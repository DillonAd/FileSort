package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config - Struct to hold configuration data
type Config struct {
	DestinationPath string
	FilePattern     string
}

// ParseConfiguration - Reads file and returns the parsed configuration
func ParseConfiguration(fileName string) []Config {
	var configs []Config

	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(contents, &configs)
	if err != nil {
		log.Fatalln(err)
	}

	return configs
}
