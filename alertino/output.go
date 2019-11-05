package alertino

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"alertino/models"
	"alertino/util"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

func (a *Alertino) processOutput(output *models.IOOutput, outputItem *models.QueueOutputItem) error {

	var errors []error

	log := logrus.WithFields(logrus.Fields{
		"outputItem": util.Dump(outputItem),
	})

	if output.StdOut {
		log.Info("processed output")
	}

	if output.WebHook != nil {
		body, err := json.Marshal(outputItem.Args)
		if err != nil {
			errors = append(errors, fmt.Errorf("failed to marshal args"))
			goto afterWebhook
		}

		_, err = a.httpClient.Post(*output.WebHook, binding.MIMEJSON, bytes.NewBuffer(body))
		if err != nil {
			errors = append(errors, fmt.Errorf("failed to send webhook request"))
			goto afterWebhook
		}

		log.WithField("webHook", *output.WebHook).Info("sent webhook request")
	}
afterWebhook:

	if len(errors) > 0 {
		var errorStrings []string
		for _, err := range errors {
			errorStrings = append(errorStrings, err.Error())
		}
		err := fmt.Errorf("multiple errors found:\n%s", strings.Join(errorStrings, "\n- "))
		return err
	}

	return nil
}
