package middleware

import "github.com/gin-gonic/gin"

// Jwt jwt
func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Next()
	}
}

