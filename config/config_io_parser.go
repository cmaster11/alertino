package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Merges multiple config files into one single config
func ParseIOConfigFiles(files []string) (*IOConfig, error) {

	finalConfig := &IOConfig{}

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("failed to read config file %s: %w", file, err)
		}

		singleConfig, err := parseIOConfigFromBytes(data)
		if err != nil {
			return nil, fmt.Errorf("failed to parse config for file %s: %w", file, err)
		}

		if finalConfig, err = finalConfig.Merge(singleConfig); err != nil {
			return nil, fmt.Errorf("failed to merge config for file %s: %w", file, err)
		}
	}

	return finalConfig, nil
}

func parseIOConfigFromBytes(data []byte) (*IOConfig, error) {
	config := &IOConfig{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config %w", err)
	}

	return config, config.Validate()
}
