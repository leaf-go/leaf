package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ErrorNotFound 404
func ErrorNotFound() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(404, gin.H{
			"message": "api not found",
		})
	}
}

// Cors 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求头部
		if origin := c.Request.Header.Get("Origin"); origin != "" {
			// 这是允许访问所有域
			c.Header("Access-Control-Allow-Origin", "*")

			//服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma, X-Token")
			//允许跨域设置
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			//缓存请求信息 单位为秒
			c.Header("Access-Control-Max-Age", "172800")
			//跨域请求是否需要带cookie信息 默认设置为true
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("Content-Type", "application/json")
		}

		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		//请求方法 当前没有意义 暂时去掉
		//method := c.Request.Action
		//if method == "OPTIONS" { // 放开探针
		//	探针是不是还需要向下进行?? 还是直接return
		//c.JSON(http.StatusOK, "Options Request!")
		//}
		// 处理请求
		c.Next() //  处理请求
	}
}

// Catch 捕获错误
func Catch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
