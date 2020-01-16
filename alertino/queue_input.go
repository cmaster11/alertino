package alertino

import (
	"github.com/cmaster11/alertino/features/io"
)

func (a *Alertino) queueProcessInputItem(inputItem *io.QueueInputItem) {

	// Check if there is any rule matching the input. If it exists, forwards to desired output
	for _, rule := range a.IOConfig.Rules {
		if rule.MatchesQueueInputItem(inputItem) {
			for _, outputId := range rule.OutputIds {
				outputItem := &io.QueueOutputItem{
					OutputId:     outputId,
					MatchedRules: []string{rule.Tag},
					Args:         inputItem.Args,
				}

				a.queueProcessOutputItem(outputItem)
			}
		}
	}

}
