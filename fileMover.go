package main

import (
	"log"
	"os"
)

type FileMover interface {
	MoveFile(string)
}

type fileMover struct {
	config []Config
}

func NewFileMover(configs []Config) FileMover {
	fm := &fileMover{}
	fm.config = configs

	return fm
}

func (fm *fileMover) MoveFile(fileName string) {
	//source := transit.sourceDirectory + transit.fileName
	//destination := transit.destinationDirectory + transit.fileName
	var source string
	var destination string
	//TODO match to config
	err := os.Rename(source, destination)
	if err != nil {
		log.Fatalln(err)
	}
}
