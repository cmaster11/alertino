package models

import (
	"alertino/util"
)

// We want to be able to parse multiple incoming payloads depending on the input, e.g. Grafana alert.
// Here's the input configuration
type IOInput struct {
	util.Validable

	// Contains the go-template used to calculate the hash of the incoming payload
	HashTemplate *util.Template `yaml:"hashTemplate"` // validate:"required"
}
