package mounts

import (
	"github.com/leaf-go/daemon"
	"github.com/leaf-go/x"
)

type TaskApplication struct {
	handler *daemon.DefaultDaemonManager
}

func NewTaskApplication() x.IApplication {
	return &TaskApplication{handler: daemon.DaemonManager}
}

func (t TaskApplication) Boot() error {
	t.handler.Start()
	return nil
}

func (t TaskApplication) Handler() interface{} {
	return t.handler
}

func (t TaskApplication) Bootstrap(boots ...x.IBootstrap) x.IApplication {
	for _, boot := range boots {
		boot.Boot(t)
	}

	return t
}

func (t TaskApplication) AutoConfig() x.IApplication {
	panic("implement me")
}

func (t TaskApplication) Shutdown() {
	t.handler.Stop()
}

func (t TaskApplication) Config(config interface{}) x.IApplication {
	panic("implement me")
}
