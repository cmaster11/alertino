package models

import (
	"alertino/util"
)

// QueueInputItem holds a processed input result
type QueueInputItem struct {
	util.Validable

	InputId string

	Args map[string]interface{}

	Hash *string
}
