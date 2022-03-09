package middleware

import "github.com/gin-gonic/gin"

// ErrorNotFound 404
func ErrorNotFound() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(404, gin.H{
			"message": "api not found",
		})
	}
}
