package models

import (
	"alertino/util"
)

type AlertRuleCondition struct {
	util.Validable

	// The input to process with this condition
	InputId string `yaml:"inputId" validate:"required"`

	// The conditions. Treat as inner blocks of a go-template if, e.g. eq .name "hello"
	If []*util.IfTemplate `yaml:"if"`
}
