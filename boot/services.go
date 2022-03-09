package boot

import (
	"github.com/gin-gonic/gin"
	"leaf-go/apps/task"
	"leaf-go/mounts"
	"leaf-go/routes"
	"x"
)

func registerServices() {
	x.Register("http.api", func() x.IApplication {
		return x.NewHttp(gin.New()).Config(Configs.Http).Bootstrap(
			// 加载路由
			routes.API{},
		)
	})

	x.Register("http.task", func() x.IApplication {
		return mounts.NewTaskApplication().Bootstrap(
			// 加载路由
			task.APITask{},
		)
	})
}

