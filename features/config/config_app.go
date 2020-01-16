package config

import (
	"github.com/cmaster11/alertino/platform/util"
)

// Represents the configuration for the app
type AppConfig struct {
	util.Validable

	ListenAddr *string `yaml:"listenAddr"`

	// ArangoDb settings
	ArangoDB *ArangoDBConfiguration `yaml:"arangoDB" validate:"required"`

	// Authentication
	OAuth2Config *struct {
		ClientId     string `yaml:"clientId" validate:"required"`
		ClientSecret string `yaml:"clientSecret" validate:"required"`
		RedirectURL  string `yaml:"redirectUrl" validate:"required,url"`
	} `yaml:"oAuth2Config" validate:"required"`

	SessionSecret string `yaml:"sessionSecret" validate:"required"`
}

type ArangoDBConfiguration struct {
	Coordinators []ArangoDBCoordinatorConfiguration `yaml:"coordinators"`

	Database string `yaml:"database" validate:"required"`
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
}

type ArangoDBCoordinatorConfiguration struct {
	Host string `yaml:"host" validate:"required"`
	Port uint   `yaml:"port" validate:"required"`
}
