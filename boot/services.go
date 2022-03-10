package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/leaf-go/x"
	"leaf/apps/task"
	"leaf/mounts"
	"leaf/routes"
	
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

