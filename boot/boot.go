package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/leaf-go/x"
	"leaf/mounts"
	"time"
	
)

var (
	Configs *mounts.Configs
)

// 初始化
func init() {
	x.SetEnv(gin.ReleaseMode, func() string {
		return gin.Mode()
	})

	// 初始化时间为当前时区
	local, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = local

	Configs = &mounts.Configs{}
	if err := Configs.Parse("./config"); err != nil {
		panic(err)
	}
}

// Application 应用
func Application() {
	// 初始化数据库、日志
	Configs.Initialize()
	// 注册app
	registerServices()
}

// Script 脚本
func Script() {
	// 初始化数据库、日志
	Configs.Initialize()
}
