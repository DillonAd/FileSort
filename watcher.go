package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/go-fsnotify/fsnotify"
)

//Watcher - Watcher for changes to file system
type Watcher interface {
	AddWatcherDirectory(string)
	GetWatchedDirectories() []string
	Start()
	Stop()
}

type fswatcher struct {
	watcher     fsnotify.Watcher
	directories []string
	endure      bool
}

// NewWatcher - Creates a new instance of Watcher
func NewWatcher() Watcher {
	return new(fswatcher)
}

func (w *fswatcher) GetWatchedDirectories() []string {
	return w.directories
}

func (w *fswatcher) Start() {
	go start(w)
}

func (w *fswatcher) Stop() {
	w.watcher.Close()
}

// AddWatcherDirectory - Adds directory to be watched and recursively adds all directories within that directory
func (w *fswatcher) AddWatcherDirectory(directory string) {
	addWatcherDirectory(w, directory, 0)
}

func start(w *fswatcher) {
	done := make(chan bool)

	go func() {
		for w.endure {
			select {
			case event := <-w.watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					//fileChanged(event.Name)
				}
			case err := <-w.watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	<-done
}

func addWatcherDirectory(watcher *fswatcher, directory string, depth int) {
	err := watcher.watcher.Add(directory)
	CheckError(err)

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
	CheckError(err)

	for _, f := range files {
		if f.IsDir() {
			fileInfo = append(fileInfo, f)
		}
	}

	return fileInfo
}
