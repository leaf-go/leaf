package middleware

import (
	"github.com/gin-gonic/gin"
	"x"
)

func Catch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer x.Recover(func(r interface{}, message string) {


		})
	}
}
