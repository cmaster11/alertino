package main

import (
	"alertino/alertino"
	"alertino/config"
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	"os"
)

var opts struct {
	config.AppConfig

	LogLevel *string `short:"v" long:"logLevel" description:"Literal log level"`

	ConfigFiles []string `short:"c" long:"config" required:"true" description:"One or multiple input/output configuration files to use"`
}

func main() {
	_, err := flags.Parse(&opts)

	if err != nil {
		//logrus.Fatalf("failed to parse flags: %s", err)
		os.Exit(1)
	}

	if opts.LogLevel != nil {
		logLevel, err := logrus.ParseLevel(*opts.LogLevel)
		if err != nil {
			logrus.Fatalf("failed to parse log level: %s", err)
		}
		logrus.SetLevel(logLevel)
	}

	ioConfig, err := config.ParseIOConfigFiles(opts.ConfigFiles)
	if err != nil {
		logrus.Fatalf("failed to parse io config: %s", err)
	}

	logrus.WithFields(logrus.Fields{
		"appConfig": opts.AppConfig,
		"ioConfig":  ioConfig,
	}).Debug("loaded configs")

	instance := alertino.Alertino{
		AppConfig: &opts.AppConfig,
		IOConfig:  ioConfig,
	}

	instance.Run()
}
