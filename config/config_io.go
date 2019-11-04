package config

import (
	"alertino/util"
	"fmt"
)

// Represents the configuration for incoming alerts and outgoing things
type IOConfig struct {
	util.Validable

	// Incoming sources, e.g. `grafana-prod`
	Sources map[string]*InputSource `yaml:"sources" validate:"dive,keys,urlfriendly,endkeys,required"`
}

// We want to be able to parse multiple incoming payloads depending on the source, e.g. Grafana alert.
// Here's the source configuration
type InputSource struct {
	util.Validable

	// Contains the go-template used to calculate the hash of an alert
	HashTemplate *util.Template `yaml:"hashTemplate"` // validate:"required"
}

// Merges current config by importing another config. Returns current config after merge
func (c *IOConfig) Merge(otherConfig *IOConfig) (*IOConfig, error) {

	if c.Sources == nil {
		c.Sources = make(map[string]*InputSource)
	}

	// Merge sources
	for key, src := range otherConfig.Sources {
		if _, found := c.Sources[key]; found {
			return nil, fmt.Errorf("config source %s already exists", key)
		}

		c.Sources[key] = src
	}

	return c, nil
}
