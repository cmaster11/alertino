package alerts

import (
	"fmt"

	"github.com/cmaster11/alertino/features/io"
	"github.com/cmaster11/alertino/platform/util"
	"github.com/sirupsen/logrus"
)

// A rule defines which inputs are bound to which outputs, and the logic that comes with this binding
type AlertRule struct {
	util.Validable

	// A friendly tag/label to identify the rule
	Tag string `yaml:"tag" validate:"required"`

	// If an alert rises, to which outputs should it be sent
	OutputIds []string `yaml:"outputIds" validate:"required"`

	// The conditions, which will trigger an alert
	When []*AlertRuleCondition `yaml:"when"`
}

func (r *AlertRule) Validate() error {
	if err := util.Validate.Struct(r); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	// TODO (05/11/2019 - cmaster11): Support multiple inputs
	var lastInputId string
	for _, when := range r.When {
		if lastInputId == "" {
			lastInputId = when.InputId
			continue
		}

		if lastInputId != when.InputId {
			return fmt.Errorf("for now only one input at a time is supported for one rule, found multiple at: %s", r.Tag)
		}
	}

	return nil
}

func (r *AlertRule) MatchesQueueInputItem(item *io.QueueInputItem) bool {

	match := true

	for _, condition := range r.When {
		if condition.InputId != item.InputId {
			match = false
			break
		}

		for _, conditionIf := range condition.If {
			conditionMatch, err := conditionIf.Matches(item.Args)
			if err != nil {
				match = false
				logrus.WithFields(logrus.Fields{
					"ruleTag": r.Tag,
					"inputId": item.InputId,
					"args":    item.Args,
				}).Errorf("error matching condition: %s", err)
				break
			}

			if !conditionMatch {
				match = false
				break
			}
		}
	}

	return match
}
