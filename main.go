package main

import (
	"flag"
)

func main() {
	var configFile string

	//TODO Expand description to describe the necessary JSON format
	flag.StringVar(&configFile, "config", "", "Configuration File")
}
