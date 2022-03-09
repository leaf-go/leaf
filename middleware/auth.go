package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

//func getToken(ctx *gin.Context) (token string, err error)  {
	//token := ctx.GetHeader("Authorization")
//}
