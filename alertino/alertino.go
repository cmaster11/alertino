package alertino

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/cmaster11/alertino/features/config"
	"github.com/cmaster11/alertino/platform/util"
	"github.com/coreos/go-oidc"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type Alertino struct {
	AppConfig *config.AppConfig
	IOConfig  *config.IOConfig

	// DB
	ArangoDBCollections *ArangoDBCollections

	// Internals
	httpClient   *http.Client
	oidcProvider *oidc.Provider
	oidcVerifier *oidc.IDTokenVerifier
	oAuth2Config *oauth2.Config
	sessionStore cookie.Store
}

func (a *Alertino) Run() {

	a.AppConfig.PanicIfInvalid()
	a.IOConfig.PanicIfInvalid()

	// --- Initialize internals
	a.ArangoDBCollections = setupArangoDB(a.AppConfig.ArangoDB)

	// Output
	a.httpClient = &http.Client{}

	// Mount all input apis
	router := util.NewRouter()
	a.setupSessions(router)
	// a.setupOAuth2(router)
	a.mountRoutes(router)

	// Start the main app loop
	go a.loop()

	listenAddress := ":8080"
	if a.AppConfig.ListenAddr != nil {
		listenAddress = *a.AppConfig.ListenAddr
	}
	util.PanicIfError(router.Run(listenAddress))
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
