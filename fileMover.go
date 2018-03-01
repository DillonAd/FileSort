package main

import (
	"os"
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
	//source := transit.sourceDirectory + transit.fileName
	//destination := transit.destinationDirectory + transit.fileName
	var source string
	var destination string

	//TODO Actually move the file to
	return os.Rename(source, destination)
}
