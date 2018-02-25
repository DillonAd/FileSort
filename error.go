package main

import "log"

// CheckError - Checks for an error and logs a fatal message if
func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
