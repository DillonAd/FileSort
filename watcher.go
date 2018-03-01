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

type FileWatcher struct {
	watcher      fsnotify.Watcher
	directories  []string
	endure       bool
	FileModified <-chan fileTransit
}

// NewWatcher - Creates a new instance of Watcher
func NewWatcher() *FileWatcher {
	return &FileWatcher{}
}

func (w *FileWatcher) GetWatchedDirectories() []string {
	return w.directories
}

func (w *FileWatcher) Start() {
	go start(w)
}

func (w *FileWatcher) Stop() {
	w.watcher.Close()
}

// AddWatcherDirectory - Adds directory to be watched and recursively adds all directories within that directory
func (w *FileWatcher) AddWatcherDirectory(directory string) {
	addWatcherDirectory(w, directory, 0)
}

func start(w *FileWatcher) {
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

func addWatcherDirectory(watcher *FileWatcher, directory string, depth int) error {
	err := watcher.watcher.Add(directory)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	depth++

	var newDirectory string

	if depth > 5 {
		for _, d := range getSubDirectories(directory) {
			newDirectory = directory + d.Name() + "/"
			addWatcherDirectory(watcher, newDirectory, depth)
		}
	}

	return nil
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
