package alertino

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"alertino/config"
	"alertino/util"
	"github.com/coreos/go-oidc"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"golang.org/x/oauth2"
)

type Alertino struct {
	AppConfig *config.AppConfig
	IOConfig  *config.IOConfig

	// Internals
	httpClient   *http.Client
	oidcProvider *oidc.Provider
	oidcVerifier *oidc.IDTokenVerifier
	oAuth2Config *oauth2.Config
	sessionStore cookie.Store

	redisClient *redis.Client
}

func (a *Alertino) Run() {

	a.AppConfig.PanicIfInvalid()
	a.IOConfig.PanicIfInvalid()

	// --- Initialize internals
	a.setupRedis()

	// Output
	a.httpClient = &http.Client{}

	// Mount all input apis
	router := util.NewRouter()
	a.setupSessions(router)
	a.setupOAuth2(router)
	a.mountRoutes(router)

	listenAddress := ":8080"
	if a.AppConfig.ListenAddr != nil {
		listenAddress = *a.AppConfig.ListenAddr
	}
	util.PanicIfError(router.Run(listenAddress))
}

func (a *Alertino) setupRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     a.AppConfig.RedisConfig.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	a.redisClient = client
	_, err := client.Ping().Result()
	util.PanicIfError(err)
}

func (a *Alertino) setupSessions(router *gin.Engine) {
	sessionSecretBytes, err := base64.StdEncoding.DecodeString(a.AppConfig.SessionSecret)
	util.PanicIfError(err)

	if len(sessionSecretBytes) != 64 {
		util.PanicIfError(fmt.Errorf("session secret must be 64 bytes long after base64 decoding"))
	}

	a.sessionStore = cookie.NewStore(sessionSecretBytes)

	router.Use(sessions.Sessions("alertino", a.sessionStore))
}
