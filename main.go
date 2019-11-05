package main

import (
	"os"

	"alertino/alertino"
	"alertino/config"
	"alertino/util"
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
)

var opts struct {
	config.AppConfig

	LogLevel *string `short:"v" long:"logLevel" description:"Literal log level"`

	AppConfigFile string   `short:"c" long:"app-config" required:"true" description:"App configuration file to use"`
	IOConfigFiles []string `short:"i" long:"io-config" required:"true" description:"One or multiple input/output configuration files to use"`
}

func main() {
	_, err := flags.Parse(&opts)

	if err != nil {
		os.Exit(1)
	}

	if opts.LogLevel != nil {
		logLevel, err := logrus.ParseLevel(*opts.LogLevel)
		if err != nil {
			logrus.Fatalf("failed to parse log level: %s", err)
		}
		logrus.SetLevel(logLevel)
	}

	appConfig, err := config.ParseAppConfigFile(opts.AppConfigFile)
	if err != nil {
		logrus.Fatalf("failed to parse io config: %s", err)
	}

	ioConfig, err := config.ParseIOConfigFiles(opts.IOConfigFiles)
	if err != nil {
		logrus.Fatalf("failed to parse io config: %s", err)
	}

	logrus.WithFields(logrus.Fields{
		"appConfig": util.Dump(appConfig),
		"ioConfig":  util.Dump(ioConfig),
	}).Debug("loaded configs")

	instance := alertino.Alertino{
		AppConfig: appConfig,
		IOConfig:  ioConfig,
	}

	instance.Run()
}
