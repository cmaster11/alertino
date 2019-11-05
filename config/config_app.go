package config

import (
	"alertino/util"
)

// Represents the configuration for the app
type AppConfig struct {
	util.Validable

	ListenAddr *string `yaml:"listenAddr"`

	// MongoDB connection string
	MongoDBConnString string `yaml:"mongoDBConnString" required:"true"`
}
