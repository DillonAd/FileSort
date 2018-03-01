package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var configFile string
	var sourcePath string

	log.Println("Getting configuration JSON file")
	//TODO Expand description to describe the necessary JSON format
	flag.StringVar(&configFile, "config", "", "Configuration File")
	flag.StringVar(&sourcePath, "src", "", "Source Directory Path")
	flag.Parse()

	fmt.Println("Reading configuration : " + configFile)
	contents, err := ioutil.ReadFile(configFile)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Parsing configuration : " + configFile)
	configs, err := new(ConfigurationCollection).ParseConfiguration(contents)

	if err != nil {
		//TODO better log messages
		log.Fatalln(err)
	}

	fmt.Println("Running Service")
	runService(configs, sourcePath)
}
