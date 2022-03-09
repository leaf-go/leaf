package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"x"
)

// Init 初始化
func Init() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 绑定日志
		bindLogger(ctx)

		// 获取时间
		ctx.Set("start_at", time.Now())
	}
}

// 绑定日志
func bindLogger(ctx *gin.Context) {
	log := x.NewLogger(ctx.Request.Method, ctx.Request.RequestURI, ctx.ClientIP())
	ctx.Set("log", log)
}
