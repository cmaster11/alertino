package main

import (
	"github.com/jessevdk/go-flags"
	"log"
)

func main() {
	_, err := flags.Parse(&opts)

	if err != nil {
		log.Fatal(err)
	}
}
