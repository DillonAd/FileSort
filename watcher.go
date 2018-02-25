package main

import (
	"log"

	"github.com/go-fsnotify/fsnotify"
)

//Watcher - Watcher for changes to file system
type Watcher interface {
	GetWatchedDirectories() []string
	Start()
	Stop()
}

type fswatcher struct {
	watcher     *fsnotify.Watcher
	directories []string
	endure      bool
}

func (w *fswatcher) GetWatchedDirectories() []string {
	return w.directories
}

func (w *fswatcher) Start() {
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

func (w *fswatcher) Stop() {
	w.watcher.Close()
}
