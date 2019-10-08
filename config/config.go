package config

// Represents the configuration of the entire app
type Config struct {
	Sources map[string]*Source `yaml:"sources" validate:"dive,required"`
}

// We want to be able to parse multiple incoming payloads depending on the source, e.g. Grafana alert.
// Here's the source configuration
type Source struct {
	// Contains the go-template used to calculate the hash of an alert
	HashTemplate string `yaml:"hashTemplate" validate:"required"`
}
