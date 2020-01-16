package alertino

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/cmaster11/alertino/features/io"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

func (a *Alertino) apiRegisterInputHandler(router gin.IRouter, key string, input *io.IOInput) {
	fn := func(context *gin.Context) {
		log := logrus.WithFields(logrus.Fields{
			"inputKey": key,
		})
		if err := a.apiProcessInput(log, context, key, input); err != nil {
			logrus.Error(err)
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}

		context.Status(http.StatusOK)
	}

	router.POST(fmt.Sprintf("/%s", key), fn)
}

// Input entry point
func (a *Alertino) apiProcessInput(log logrus.FieldLogger, c *gin.Context, inputId string, ioInput *io.IOInput) error {
	payload := make(map[string]interface{})
	if err := c.ShouldBindWith(&payload, binding.JSON); err != nil {
		return err
	}
	log = log.WithField("payload", payload)

	log.Info("got request")

	var hash *string

	// Process deduplication, which is possible only if there is a hash template
	if ioInput.HashTemplate != nil {
		hashExecuted, err := ioInput.HashTemplate.Execute(payload)
		if err != nil {
			return fmt.Errorf("failed to calculate deduplication hash: %w", err)
		}

		md5Hash := md5.Sum([]byte(hashExecuted))
		hashString := fmt.Sprintf("%x", md5Hash)
		hash = &hashString

		log = log.WithField("hash", hashString)
		log.Debug("with hash")
	}

	queueItem := &io.QueueInputItem{
		InputId: inputId,
		Args:    payload,
		Hash:    hash,
	}

	a.queueProcessInputItem(queueItem)

	return nil
}
