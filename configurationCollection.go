package main

import (
	"encoding/json"
	"regexp"
)

//ConfigurationCollection - Collection of configurations
type ConfigurationCollection struct {
	configurations []Config
}

// ParseConfiguration - Reads file and returns the parsed configuration
func (cc *ConfigurationCollection) ParseConfiguration(contents []byte) (ConfigurationCollection, error) {
	var configs []Config
	var err error

	err = json.Unmarshal(contents, &configs)

	return ConfigurationCollection{configs}, err
}

//GetMatchingConfigurations - Returns the all configurations to match the file name
func (cc *ConfigurationCollection) GetMatchingConfigurations(fileName string) ([]Config, error) {
	var configs []Config
	var err error

	for _, cfg := range cc.configurations {
		matched, err := regexp.MatchString(cfg.FilePattern, fileName)

		if matched {
			configs = append(configs, cfg)
		}

		if err == nil {
		}
	}

	return configs, err
}
