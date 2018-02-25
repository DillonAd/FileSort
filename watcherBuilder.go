package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/go-fsnotify/fsnotify"
)

//WatcherBuilder - Fluent interface to set up Watcher
type WatcherBuilder interface {
	CreateNewWatcher() WatcherBuilder
	AddWatcherDirectory(string) WatcherBuilder
	Build() Watcher
}

type watcherBuilder struct {
	watcher fswatcher
}

// CreateNewWatcher - Creates the Watcher
func (wb *watcherBuilder) CreateNewWatcher() WatcherBuilder {
	newWatcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatalln(err)
	}

	newFsWatcher := &fswatcher{}
	newFsWatcher.watcher = newWatcher
	wb.watcher = *newFsWatcher

	return wb
}

// AddWatcherDirectory - Adds directory to be watched and recursively adds all directories within that directory
func (wb *watcherBuilder) AddWatcherDirectory(directory string) WatcherBuilder {
	addWatcherDirectory(wb.watcher, directory, 0)
	return wb
}

//Build - Finishes building the Watcher
func (wb *watcherBuilder) Build() Watcher {
	return &fswatcher{
		directories: wb.watcher.directories,
		endure:      wb.watcher.endure,
		watcher:     wb.watcher.watcher,
	}
}

func addWatcherDirectory(watcher fswatcher, directory string, depth int) {
	err := watcher.watcher.Add(directory)
	if err != nil {
		log.Fatalln(err)
	}

	depth++

	var newDirectory string

	if depth > 5 {
		for _, d := range getSubDirectories(directory) {
			newDirectory = directory + d.Name() + "/"
			addWatcherDirectory(watcher, newDirectory, depth)
		}
	}
}

func getSubDirectories(directory string) []os.FileInfo {
	var fileInfo []os.FileInfo

	files, err := ioutil.ReadDir(directory)

	if err != nil {
		log.Fatalln(err)
	}

	for _, f := range files {
		if f.IsDir() {
			fileInfo = append(fileInfo, f)
		}
	}

	return fileInfo
}
