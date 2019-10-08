package main

var opts struct {
	ConfigFiles []string `short:"c" long:"config-file" required:"true" description:"One or multiple configuration files to use"`
}
