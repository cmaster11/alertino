package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Merges multiple config files into one single config
func ParseAppConfigFromFile(file string) (*AppConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", file, err)
	}

	config, err := parseAppConfigFromBytes(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config for file %s: %w", file, err)
	}

	return config, nil
}

func parseAppConfigFromBytes(data []byte) (*AppConfig, error) {
	config := &AppConfig{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config %w", err)
	}

	return config, config.Validate()
}
