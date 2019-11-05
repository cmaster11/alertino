package alertino

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"alertino/util"
	"github.com/coreos/go-oidc"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func (a *Alertino) setupOAuth2(router *gin.Engine) {
	oidcProvider, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
	util.PanicIfError(err)

	a.oidcProvider = oidcProvider
	a.oidcVerifier = oidcProvider.Verifier(&oidc.Config{ClientID: a.AppConfig.OAuth2Config.ClientId})

	// Configure an OpenID Connect aware OAuth2 client.
	a.oAuth2Config = &oauth2.Config{
		ClientID:     a.AppConfig.OAuth2Config.ClientId,
		ClientSecret: a.AppConfig.OAuth2Config.ClientSecret,
		RedirectURL:  a.AppConfig.OAuth2Config.RedirectURL,

		// Discovery returns the OAuth2 endpoints.
		Endpoint: oidcProvider.Endpoint(),

		Scopes: []string{oidc.ScopeOpenID, "email"},
	}

	// OAuth2
	groupAuth := router.Group("/auth", a.routeHandlerRedirectIfAuthenticated)
	{
		groupAuth.GET("/oauth2/google", a.routeOAuth2LoginGoogle)
		groupAuth.GET("/oauth2/google/callback", a.routeOAuth2CallbackGoogle)
	}

	// TODO (05/11/2019 - cmaster11):
	// TODO (05/11/2019 - cmaster11): test session
	// TODO (05/11/2019 - cmaster11):
}

func (a *Alertino) routeHandlerRedirectIfAuthenticated(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get(authSessionEmailKey) != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.Next()
}

const authSessionEmailKey = "email"
const oAuth2StatePrefix = "oauth2-state-"

func (a *Alertino) routeOAuth2LoginGoogle(c *gin.Context) {
	// Generate a one-usage state
	state, err := util.NewSeqMD5Id()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to generate random state id: %w", err))
		return
	}

	now := time.Now()
	_, err = a.redisClient.Set(oAuth2StatePrefix+state, now.UnixNano(), 5*time.Minute).Result()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to store random state id: %w", err))
		return
	}

	c.Redirect(http.StatusFound, a.oAuth2Config.AuthCodeURL(state))
}

func (a *Alertino) routeOAuth2CallbackGoogle(c *gin.Context) {
	state := c.Query("state")
	// Verify that a state exists
	_, err := a.redisClient.Get(oAuth2StatePrefix + state).Result()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get oauth2 state: %w", err))
		return
	}

	// Verify state and errors.
	oauth2Token, err := a.oAuth2Config.Exchange(context.Background(), c.Query("code"))
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to exchange code: %w", err))
		return
	}

	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		_ = c.AbortWithError(http.StatusBadRequest, fmt.Errorf("missing id_token"))
		return
	}

	// Parse and verify ID Token payload.
	idToken, err := a.oidcVerifier.Verify(context.Background(), rawIDToken)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to verify id token payload: %w", err))
		return
	}

	// Extract custom claims
	var claims struct {
		Email string `json:"email"`
	}
	if err := idToken.Claims(&claims); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to extract claims: %w", err))
		return
	}

	// Save the session state
	session := sessions.Default(c)
	session.Set(authSessionEmailKey, claims.Email)
	if err := session.Save(); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to save session: %w", err))
		return
	}

	// Login was successful, so delete the state
	_, _ = a.redisClient.Del(oAuth2StatePrefix + state).Result()
}
