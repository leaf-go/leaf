package task

import (
	"github.com/leaf-go/daemon"
	"github.com/leaf-go/x"
	"fmt"
	"time"
	
)

type APITask struct {
}

func (A APITask) Boot(app x.IApplication) {
	handler := app.Handler()
	worker := handler.(*daemon.DefaultDaemonManager)

	worker.RegisterLoop("testing-loop", func() daemon.TaskFunc {
		return func(handler daemon.Handler) {
			fmt.Println("looping test loop")
		}
	}, 10*time.Second)
	//
	worker.RegisterCrontab("testing-crontab", func() daemon.TaskFunc {
		return func(handler daemon.Handler) {
			fmt.Println("crontab test task")
		}
	}, "*/2 * * * *")
}
