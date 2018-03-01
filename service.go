package main

import (
	"fmt"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
	"golang.org/x/sys/windows/svc/eventlog"
)

//ServiceName - Name of the Service
const ServiceName string = "FileSort"

var elog debug.Log

type fileSortService struct {
	configs ConfigurationCollection
	source  string
}

//Execute - Executes a command for the service
func (fss *fileSortService) Execute(args []string, cr <-chan svc.ChangeRequest, status chan<- svc.Status) (ssec bool, errno uint32) {

	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue

	watcher := NewWatcher()
	for _, c := range fss.configs.configurations {
		watcher.AddWatcherDirectory(c.DestinationPath)
	}

	watcher.Start()
	defer watcher.Stop()

	fileMover := NewFileMover(fss.source)

loop:
	for {
		select {
		case fn := <-watcher.FileModified:
			go func() {
				configs, err := fss.configs.GetMatchingConfigurations(fn)

				if err != nil {
					elog.Error(1, fmt.Sprintf("unexpected error while matching file name to configurations : #%d\n", err))
					return
				}

				if configs != nil {
					for _, cfg := range configs {
						err := fileMover.MoveFile(fn, cfg)

						if err != nil {
							elog.Error(1, fmt.Sprintf("unexpected error while moving file to destination : #%d\n", err))
						}
					}
				}
			}()
		case req := <-cr:
			switch req.Cmd {
			case svc.Interrogate:
				status <- req.CurrentStatus
			case svc.Stop, svc.Shutdown:
				break loop
			case svc.Pause:
				status <- svc.Status{State: svc.Paused, Accepts: cmdsAccepted}
			case svc.Continue:
				status <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
			default:
				elog.Error(1, fmt.Sprintf("unexpected control request #%d", req))
			}
		}
	}

	return
}

func runService(configs ConfigurationCollection, source string) {
	var err error

	elog, err = eventlog.Open(ServiceName)
	if err != nil {
		return
	}

	defer elog.Close()

	elog.Info(1, fmt.Sprintf("starting %s service", ServiceName))
	run := svc.Run

	err = run(ServiceName, &fileSortService{configs, source})
	if err != nil {
		elog.Error(1, fmt.Sprintf("%s service failed: %v", ServiceName, err))
		return
	}

	elog.Info(1, fmt.Sprintf("%s service stopped", ServiceName))
}
