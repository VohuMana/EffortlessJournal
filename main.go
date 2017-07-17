package main

import (
	"flag"
	"log"
)

func main() {
	// Load configuration
	var (
		configFile = flag.String("config", "", "Name of the config file that contains the client configuration")
	)
	flag.Parse()

	_, err := LoadConfiguration(*configFile)
	if err != nil {
		log.Fatal(err)
	}

}