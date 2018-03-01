package main

import (
	"fmt"
	"testing"
)

func TestParseConfiguration_CorrectFormat(t *testing.T) {
	const destinationPath string = "/path/to/dest"
	const filePattern string = "*.*"

	configStr := fmt.Sprintf("[{ \"DestinationPath\": \"%s\", \"FilePattern\": \"%s\" }]", destinationPath, filePattern)
	configs, err := ParseConfiguration([]byte(configStr))

	if err != nil || len(configs.configurations) != 1 {
		t.Errorf("either error was thrown : %v\n or collection had too many or too few elements : length = %v",
			err, len(configs.configurations))
	}
}

func TestParseConfiguration_IncorrectFormat(t *testing.T) {
	const destinationPath string = "/path/to/dest"
	const filePattern string = "*.*"

	configStr := fmt.Sprintf("{ \"DestinationPath\": \"%s\", \"FilePattern\": \"%s\" }", destinationPath, filePattern)
	configs, err := ParseConfiguration([]byte(configStr))

	if err == nil || len(configs.configurations) > 0 {
		t.Errorf("error not raised")
	}
}

func TestParseConfiguration_IsMatch(t *testing.T) {
	const destinationPath string = "/path/to/dest"
	const filePattern string = "[a-z]{1,4}.go"

	config := &Config{destinationPath, filePattern}

	isMatch, err := config.IsMatch("test.go")

	if err != nil || !isMatch {
		t.Errorf("either error was thrown : %v\n or file name did not match correctly (isMatch : %v)",
			err, isMatch)
	}
}

func TestParseConfiguration_IsNotMatch(t *testing.T) {
	const destinationPath string = "/path/to/dest"
	const filePattern string = "*.go"

	config := &Config{destinationPath, filePattern}

	isMatch, err := config.IsMatch("gopher.sql")

	if isMatch && (err != nil || err == nil) {
		t.Errorf("error was nil or match value was true (match value : %v)", isMatch)
	}
}
