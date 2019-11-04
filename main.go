package main

import (
	"alertino/alertino"
	"alertino/config"
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	"os"
)

var opts struct {
	LogLevel *string `short:"v" long:"log-level" description:"Literal log level"`

	ConfigFile       string   `short:"c" long:"app-config" required:"true" description:"App configuration file"`
	InputConfigFiles []string `short:"i" long:"io-config" required:"true" description:"One or multiple input/output configuration files to use"`
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

	appConfig, err := config.ParseAppConfigFromFile(opts.ConfigFile)
	if err != nil {
		logrus.Fatalf("failed to parse app config: %s", err)
	}

	ioConfig, err := config.ParseIOConfigFiles(opts.InputConfigFiles)
	if err != nil {
		logrus.Fatalf("failed to parse io config: %s", err)
	}

	logrus.WithFields(logrus.Fields{
		"appConfig": appConfig,
		"ioConfig":  ioConfig,
	}).Debug("loaded configs")

	alertino := alertino.Alertino{
		AppConfig: appConfig,
		IOConfig:  ioConfig,
	}

	alertino.Run()
}
