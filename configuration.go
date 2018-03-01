package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"regexp"
)

// Config - Struct to hold configuration data
type Config struct {
	DestinationPath string
	FilePattern     string
}

//IsMatch - Matches file to configuration's file pattern
func (c *Config) IsMatch(fileName string) (bool, error) {
	return regexp.MatchString(c.FilePattern, fileName)
}

// ParseConfiguration - Reads file and returns the parsed configuration
func ParseConfiguration(contents []byte) (ConfigurationCollection, error) {
	var configs []Config
	var err error

	err = json.Unmarshal(contents, &configs)

	return ConfigurationCollection{configs}, err
}

//ReadConfiguration - Reads contents of a file
func ReadConfiguration(fileName string) []byte {
	contents, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatalln(err)
	}

	return contents
}
