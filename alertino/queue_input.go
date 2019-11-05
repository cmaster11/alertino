package alertino

import (
	"alertino/models"
)

func (a *Alertino) queueProcessInputItem(inputItem *models.QueueInputItem) {

	// Check if there is any rule matching the input. If it exists, forwards to desired output
	for _, rule := range a.IOConfig.Rules {
		if rule.MatchesQueueInputItem(inputItem) {
			for _, outputId := range rule.OutputIds {
				outputItem := &models.QueueOutputItem{
					OutputId:     outputId,
					MatchedRules: []string{rule.Tag},
					Args:         inputItem.Args,
				}

				a.queueProcessOutputItem(outputItem)
			}
		}
	}

}
