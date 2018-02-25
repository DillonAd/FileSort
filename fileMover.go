package main

import (
	"os"
)

func moveFile(transit fileTransit) {
	source := transit.sourceDirectory
	destination := transit.destinationDirectory

	err := os.Rename(source, destination)
	CheckError(err)
}
