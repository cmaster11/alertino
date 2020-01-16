package io

import (
	"github.com/cmaster11/alertino/platform/util"
)

type QueueOutputItem struct {
	util.Validable

	OutputId string

	MatchedRules []string

	Args map[string]interface{}
}
