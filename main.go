package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var configFile string

	log.Println("Getting configuration JSON file")
	//TODO Expand description to describe the necessary JSON format
	flag.StringVar(&configFile, "config", "", "Configuration File")
	flag.Parse()
	fmt.Println("Reading configuration : " + configFile)

	configs := ParseConfiguration(configFile)
	watcher := NewWatcher()

	for _, c := range configs {
		watcher.AddWatcherDirectory(c.DestinationPath)
	}

	watcher.Start()
}
