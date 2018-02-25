package main

import (
	"encoding/json"
	"io/ioutil"
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
	CheckError(err)

	err = json.Unmarshal(contents, &configs)
	CheckError(err)

	return configs
}
