package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Merges multiple config files into one single config
func ParseConfigFiles(files []string) (*Config, error) {

	finalConfig := &Config{}

	for _, file := range files {
		singleConfig := &Config{}

		data, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, errors.Unwrap()

		}
		content, err := ioutil.ReadAll()

		if err := yaml.Unmarshal(); err != nil {

		}
	}

}
