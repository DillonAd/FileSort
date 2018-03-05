package main

import (
	"os"
	"strings"
)

//FileMover - Contract for moving files
type FileMover interface {
	MoveFile(string, Config) error
}

type fileMover struct {
	source string
}

//NewFileMover - Returns a new instance of FileMover
func NewFileMover(defaultSource string) FileMover {
	fm := &fileMover{}
	fm.source = defaultSource

	return fm
}

func (fm *fileMover) MoveFile(fileName string, config Config) error {
	source := completePath(fm.source, fileName)
	destination := completePath(config.DestinationPath, fileName)

	return os.Rename(source, destination)
}

func completePath(initPath string, fileName string) string {
	if !strings.HasSuffix(initPath, "/") {
		initPath = initPath + "/"
	}

	return initPath + fileName
}
