package task

import (
	"fmt"
	"leaf-go/mounts"
	"time"
	"x"
)

type APITask struct {
}

func (A APITask) Boot(app x.IApplication) {
	handler := app.Handler()
	worker := handler.(*mounts.DefaultDaemonManager)

	worker.RegisterLoop("testing-loop", func() mounts.TaskFunc {
		return func(handler mounts.Handler) {
			fmt.Println("looping test loop")
		}
	}, 10*time.Second)
	//
	worker.RegisterCrontab("testing-crontab", func() mounts.TaskFunc {
		return func(handler mounts.Handler) {
			fmt.Println("crontab test task")
		}
	}, "*/2 * * * *")
}
