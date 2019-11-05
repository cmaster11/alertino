package alertino

import (
	"net/http"

	"alertino/config"
	"alertino/util"
	"github.com/gin-gonic/gin"
)

type Alertino struct {
	AppConfig *config.AppConfig
	IOConfig  *config.IOConfig

	// Internals
	httpClient *http.Client
}

func (a *Alertino) Run() {

	a.AppConfig.PanicIfInvalid()
	a.IOConfig.PanicIfInvalid()

	// Initialize internals
	a.httpClient = &http.Client{}

	// Mount all input apis
	router := util.NewRouter()
	a.mountRoutes(router)

	listenAddress := ":8080"
	if a.AppConfig.ListenAddr != nil {
		listenAddress = *a.AppConfig.ListenAddr
	}
	util.PanicIfError(router.Run(listenAddress))
}

// Mount all API routes
func (a *Alertino) mountRoutes(router *gin.Engine) {

	groupInput := router.Group("/input")
	{
		for key, input := range a.IOConfig.Inputs {
			a.apiRegisterInputHandler(groupInput, key, input)
		}
	}

}
