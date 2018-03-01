package main

import (
	"fmt"
	"testing"
)

func TestParseConfiguration_CorrectFormat(t *testing.T) {
	const destinationPath string = "/path/to/dest"
	const filePattern string = "*.*"

	configStr := fmt.Sprintf("[{ \"DestinationPath\": \"%s\", \"FilePattern\": \"%s\" }]", destinationPath, filePattern)
	configs, err := new(ConfigurationCollection).ParseConfiguration([]byte(configStr))

	if err != nil || len(configs.configurations) != 1 {
		t.Errorf("either error was thrown : %v\n or collection had too many or too few elements : length = %v",
			err, len(configs.configurations))
	}
}

func TestParseConfiguration_IncorrectFormat(t *testing.T) {
	const destinationPath string = "/path/to/dest"
	const filePattern string = "*.*"

	configStr := fmt.Sprintf("{ \"DestinationPath\": \"%s\", \"FilePattern\": \"%s\" }", destinationPath, filePattern)
	configs, err := new(ConfigurationCollection).ParseConfiguration([]byte(configStr))

	if err == nil || len(configs.configurations) > 0 {
		t.Errorf("error not raised")
	}
}
