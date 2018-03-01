package main

import (
	"os"
	"regexp"
)

//FileMover - Contract for moving files
type FileMover interface {
	MatchFile(string) (bool, Config, error)
	MoveFile(string, Config) error
}

type fileMover struct {
	configs []Config
	source  string
}

//NewFileMover - Returns a new instance of FileMover
func NewFileMover(configs []Config, defaultSource string) FileMover {
	fm := &fileMover{}
	fm.configs = configs
	fm.source = defaultSource
	return fm
}

func (fm *fileMover) MatchFile(fileName string) (bool, Config, error) {

	var cfg Config
	var err error
	matched := false

	for _, cfg = range fm.configs {
		matched, err := regexp.MatchString(cfg.FilePattern, fileName)

		if err == nil && matched {
			break
		}
	}

	return matched, cfg, err
}

func (fm *fileMover) MoveFile(fileName string, config Config) error {
	//source := transit.sourceDirectory + transit.fileName
	//destination := transit.destinationDirectory + transit.fileName
	var source string
	var destination string

	//TODO Actually move the file to
	err := os.Rename(source, destination)

	return err
}
