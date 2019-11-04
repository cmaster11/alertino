package alertino

import (
	"alertino/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (a *Alertino) registerInputHandler(router gin.IRouter, key string, source *config.InputSource) {
	fn := func(context *gin.Context) {
		log := logrus.WithFields(logrus.Fields{
			"inputKey": key,
		})
		if err := a.processInput(log, context, key, source); err != nil {
			logrus.Error(err)
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}

		context.Status(http.StatusOK)
	}

	router.POST(fmt.Sprintf("/%s", key), fn)
}

// Input entry point
func (a *Alertino) processInput(log logrus.FieldLogger, c *gin.Context, key string, source *config.InputSource) error {
	payload := make(map[string]interface{})
	if err := c.ShouldBindWith(&payload, binding.JSON); err != nil {
		return err
	}
	log = log.WithField("payload", payload)

	log.Info("got request")

	// Process deduplication, which is possible only if there is a hash template
	if source.HashTemplate != nil {
		hashTemplate, err := source.HashTemplate.Execute(payload)
		if err != nil {
			return fmt.Errorf("failed to calculate deduplication hash: %w", err)
		}

		log = log.WithField("hashTemplate", hashTemplate)
		log.Debug("with hash")
	}

	return nil
}
