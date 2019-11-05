package alertino

import (
	"alertino/models"
	"alertino/util"
	"github.com/sirupsen/logrus"
)

func (a *Alertino) queueProcessOutputItem(outputItem *models.QueueOutputItem) {

	output, found := a.IOConfig.Outputs[outputItem.OutputId]
	if !found {
		logrus.WithFields(logrus.Fields{
			"outputItem": util.Dump(outputItem),
		}).Error("output not found")
	}

	if err := a.processOutput(output, outputItem); err != nil {
		logrus.WithFields(logrus.Fields{
			"outputItem": util.Dump(outputItem),
		}).WithError(err).Error("failed to process output")
	}

}
