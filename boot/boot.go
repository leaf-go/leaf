package boot

import (
	"github.com/gin-gonic/gin"
	"leaf-go/apps/task"
	"leaf-go/mounts"
	"leaf-go/routes"
	"x"
)

func Application() {
	// 注册app
	registerServices()

	// 数据库
	database()

	// log
	log()
}

func log()  {
	x.Configs.Log.Initialize()
}

// Database 数据库启动
func database()  {
	// 初始化mysql
	x.Configs.Mysql.Init()

	// 初始化redis
	x.Configs.Redis.Init()
}


func registerServices() {
	x.Register("http.api", func() x.IApplication {
		return x.NewHttp(gin.New()).AutoConfig().Bootstrap(
			// 加载路由
			routes.APIRouter{},
		)
	})

	x.Register("http.task", func() x.IApplication {
		return mounts.NewTaskApplication().Bootstrap(
			task.APITask{},
		)
	})
}
