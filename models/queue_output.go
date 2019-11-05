package models

import (
	"alertino/util"
)

type QueueOutputItem struct {
	util.Validable

	OutputId string

	MatchedRules []string

	Args map[string]interface{}
}
