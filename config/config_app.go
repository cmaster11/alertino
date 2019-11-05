package config

import (
	"alertino/util"
)

// Represents the configuration for the app
type AppConfig struct {
	util.Validable

	ListenAddr *string `yaml:"listenAddr"`

	// MongoDB connection string
	MongoDBConnString string `yaml:"mongoDBConnString" validate:"required"`

	// Authentication
	OAuth2Config *struct {
		ClientId     string `yaml:"clientId" validate:"required"`
		ClientSecret string `yaml:"clientSecret" validate:"required"`
		RedirectURL  string `yaml:"redirectUrl" validate:"required,url"`
	} `yaml:"oAuth2Config" validate:"required"`

	RedisConfig *struct {
		Addr string `yaml:"addr" validate:"required"`
	} `yaml:"redisConfig" validate:"required"`

	SessionSecret string `yaml:"sessionSecret" validate:"required"`
}
