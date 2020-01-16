package alertino

import (
	"github.com/cmaster11/alertino/features/io"
	"github.com/cmaster11/alertino/platform/util"
	"github.com/sirupsen/logrus"
)

func (a *Alertino) queueProcessOutputItem(outputItem *io.QueueOutputItem) {

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
