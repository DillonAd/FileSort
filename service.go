package main

import (
	"fmt"
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
	"golang.org/x/sys/windows/svc/eventlog"
)

//ServiceName - Name of the Service
const ServiceName string = "FileSort"

var elog debug.Log

type fileSortService struct{}

//Execute - Executes a command for the service
func (fss *fileSortService) Execute(args []string, cr <-chan svc.ChangeRequest, status chan<- svc.Status) (ssec bool, errno uint32) {

	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue

	watcher := NewWatcher()
	//TODO get config data
	fileMover := NewFileMover(nil)

	for {
		select {
		case ft := <-watcher.FileModified:
			time.Sleep(time.Second)
			fileMover.MoveFile(ft)
		case req := <-cr:
			switch req.Cmd {
			case svc.Interrogate:
				status <- req.CurrentStatus
			case svc.Stop, svc.Shutdown:
				break
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

func runService(isDebug bool) {
	var err error
	if isDebug {
		elog = debug.New(ServiceName)
	} else {
		elog, err = eventlog.Open(ServiceName)
		if err != nil {
			return
		}
	}
	defer elog.Close()

	elog.Info(1, fmt.Sprintf("starting %s service", ServiceName))
	run := svc.Run
	if isDebug {
		run = debug.Run
	}
	err = run(ServiceName, &fileSortService{})
	if err != nil {
		elog.Error(1, fmt.Sprintf("%s service failed: %v", ServiceName, err))
		return
	}
	elog.Info(1, fmt.Sprintf("%s service stopped", ServiceName))
}
