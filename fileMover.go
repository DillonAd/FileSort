package main

import (
	"log"
	"os"
)

type FileMover interface {
	MoveFile(fileTransit)
}

type fileMover struct {
	config []Config
}

func NewFileMover(configs []Config) FileMover {
	fm := &fileMover{}
	fm.config = configs

	return fm
}

func (fm *fileMover) MoveFile(transit fileTransit) {
	source := transit.sourceDirectory
	destination := transit.destinationDirectory

	err := os.Rename(source, destination)
	if err != nil {
		log.Fatalln(err)
	}
}
