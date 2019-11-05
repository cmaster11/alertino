package models

import (
	"fmt"

	"alertino/util"
)

// We want to be able to send data to multiple outputs depending on rules
type IOOutput struct {
	util.Validable

	StdOut  bool    `yaml:"stdOut"`
	WebHook *string `yaml:"webHook" validate:"omitempty,url"`
}

func (v *IOOutput) Validate() error {
	if err := util.Validate.Struct(v); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	hasOutputMethod := false
	if v.StdOut {
		hasOutputMethod = true
	}

	if v.WebHook != nil {
		hasOutputMethod = true
	}

	if !hasOutputMethod {
		return fmt.Errorf("at least one output method must be defined")
	}

	return nil
}
