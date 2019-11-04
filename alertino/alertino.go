package alertino

import (
	"alertino/config"
	"alertino/util"
)

type Alertino struct {
	AppConfig *config.AppConfig
	IOConfig  *config.IOConfig
}

func (a *Alertino) Run() {

	a.AppConfig.PanicIfInvalid()
	a.IOConfig.PanicIfInvalid()

	// Mount all input apis

	router := util.NewRouter()

	groupInput := router.Group("/input")
	{
		for key, input := range a.IOConfig.Sources {
			a.registerInputHandler(groupInput, key, input)
		}
	}

	util.PanicIfError(router.Run())
}
