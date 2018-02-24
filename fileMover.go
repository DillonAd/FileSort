package main

import (
	"log"
	"os"
)

func moveFile(transit fileTransit) {
	source := transit.sourceDirectory
	destination := transit.destinationDirectory

	err := os.Rename(source, destination)

	if err != nil {
		log.Fatalln(err)
	}
}
