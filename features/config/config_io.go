package config

import (
	"fmt"

	"github.com/cmaster11/alertino/features/alerts"
	"github.com/cmaster11/alertino/features/io"
	"github.com/cmaster11/alertino/platform/util"
)

// Represents the configuration for incoming alerts and outgoing things
type IOConfig struct {
	util.Validable

	// Incoming inputs, e.g. `grafana-prod`
	Inputs map[string]*io.IOInput `yaml:"inputs" validate:"dive,keys,urlfriendly,endkeys,required"`

	// Outputs, e.g. `email`, `webhook`
	Outputs map[string]*io.IOOutput `yaml:"outputs" validate:"dive,keys,urlfriendly,endkeys,required"`

	// The rules which will trigger alerts generation
	Rules []*alerts.AlertRule `yaml:"rules"`
}

// Merges current config by importing another config. Returns current config after merge
func (c *IOConfig) Merge(otherConfig *IOConfig) (*IOConfig, error) {

	// Merge inputs
	if c.Inputs == nil {
		c.Inputs = make(map[string]*io.IOInput)
	}

	for key, src := range otherConfig.Inputs {
		if _, found := c.Inputs[key]; found {
			return nil, fmt.Errorf("config input %s already exists", key)
		}

		c.Inputs[key] = src
	}

	// Merge outputs
	if c.Outputs == nil {
		c.Outputs = make(map[string]*io.IOOutput)
	}
	for key, output := range otherConfig.Outputs {
		if _, found := c.Outputs[key]; found {
			return nil, fmt.Errorf("config output %s already exists", key)
		}

		c.Outputs[key] = output
	}

	// Merge rules
	for _, rule := range otherConfig.Rules {
		c.Rules = append(c.Rules, rule)
	}

	return c, nil
}

func (c *IOConfig) Validate() error {
	if err := util.Validate.Struct(c); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	// Check that there are no mistypings in the rules
	for _, rule := range c.Rules {
		for _, outputId := range rule.OutputIds {
			if _, found := c.Outputs[outputId]; !found {
				return fmt.Errorf("validation error: output %s not found", outputId)
			}
		}

		for _, condition := range rule.When {
			if _, found := c.Inputs[condition.InputId]; !found {
				return fmt.Errorf("validation error: input %s not found", condition.InputId)
			}
		}
	}

	return nil
}
