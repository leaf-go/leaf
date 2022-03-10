package middleware

import (
	"github.com/gin-gonic/gin"
	"leaf/mounts"
)

func Init() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 初始化日志
		ctx.Set("logger", mounts.Logger())



	}
}


