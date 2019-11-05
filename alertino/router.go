package alertino

import (
	"github.com/gin-gonic/gin"
)

// Mount all API routes
func (a *Alertino) mountRoutes(router *gin.Engine) {

	// Input
	groupInput := router.Group("/input")
	{
		for key, input := range a.IOConfig.Inputs {
			a.apiRegisterInputHandler(groupInput, key, input)
		}
	}

}
