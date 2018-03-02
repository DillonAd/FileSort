package main

import (
	"fmt"
	"testing"
)

func TestNewConfigurationCollection_CorrectFormat(t *testing.T) {
	const destinationPath string = "/path/to/dest"
	const filePattern string = "*.*"

	configStr := fmt.Sprintf("[{ \"DestinationPath\": \"%s\", \"FilePattern\": \"%s\" }]", destinationPath, filePattern)
	configs, err := NewConfigurationCollection([]byte(configStr))

	if err != nil || len(configs.configurations) != 1 {
		t.Errorf("either error was thrown : %v\n or collection had too many or too few elements : length = %v",
			err, len(configs.configurations))
	}
}

func TestParseConfiguration_IncorrectFormat(t *testing.T) {
	const destinationPath string = "/path/to/dest"
	const filePattern string = "*.*"

	configStr := fmt.Sprintf("{ \"DestinationPath\": \"%s\", \"FilePattern\": \"%s\" }", destinationPath, filePattern)
	configs, err := NewConfigurationCollection([]byte(configStr))

	if err == nil || len(configs.configurations) > 0 {
		t.Errorf("error not raised")
	}
}

//TODO Add second item to JSON array
func TestGetMatchingConfigurations_Single(t *testing.T) {
	const destinationPath string = "/path/to/dest"
	const filePattern string = "[a-z]{1,4}.go"

	configStr := fmt.Sprintf("[{ \"DestinationPath\": \"%s\", \"FilePattern\": \"%s\" }, { \"DestinationPath\": \"/path/to/dir1\", \"FilePattern\": \"go1.sql\" }]", destinationPath, filePattern)
	configColl, err := NewConfigurationCollection([]byte(configStr))
	configs, err := configColl.GetMatchingConfigurations("test.go")

	if err != nil && len(configColl.configurations) != 1 {
		t.Errorf("either error was thrown : %v\n or collection had too many or too few elements : length = %v",
			err, len(configs))
	}
}

func TestGetMatchingConfigurations_None(t *testing.T) {
	const destinationPath string = "/path/to/dest"
	const filePattern string = "[a-z]{1,4}.go"

	configStr := fmt.Sprintf("[{ \"DestinationPath\": \"%s\", \"FilePattern\": \"%s\" }, { \"DestinationPath\": \"/path/to/dir1\", \"FilePattern\": \"go1.sql\" }]", destinationPath, filePattern)
	configColl, err := NewConfigurationCollection([]byte(configStr))
	configs, err := configColl.GetMatchingConfigurations("gopher.sql")

	if err != nil && len(configColl.configurations) == 1 {
		t.Errorf("either error was thrown : %v\n or collection had too many or too few elements : length = %v",
			err, len(configs))
	}
}

func TestGetMatchingConfigurations_Mutliple(t *testing.T) {
	const destinationPath string = "/path/to/dest"
	const filePattern string = "[a-z]{1,4}.[a-z]{1,4}"

	configStr := fmt.Sprintf("[{ \"DestinationPath\": \"%s\", \"FilePattern\": \"%s\" }, { \"DestinationPath\": \"/path/to/dir1\", \"FilePattern\": \"go1.sql\" }]", destinationPath, filePattern)
	configColl, err := NewConfigurationCollection([]byte(configStr))
	configs, err := configColl.GetMatchingConfigurations("gopher.sql")

	if err != nil && len(configColl.configurations) != 2 {
		t.Errorf("either error was thrown : %v\n or collection had too many or too few elements : length = %v",
			err, len(configs))
	}
}
