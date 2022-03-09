package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"x"
)

func Catch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer x.Recover(func(r interface{}, message string) {

			log, _ := ctx.Get("log")
			logger := log.(x.Logger)

			result := transformCatch(logger, r)

			ctx.AbortWithStatusJSON(http.StatusOK, result)
		})

		ctx.Next()
	}
}

// 转换成响应内容
func transformCatch(logger x.Logger, r interface{}) (result gin.H) {
	switch r.(type) {
	case x.IErrno:
		errno := r.(x.IErrno)
		result = gin.H{
			"code":     errno,
			"message":  x.GetErrorMessage(errno),
			"trace_id": logger.GetId(),
		}
		break
	case *x.Error:
		e := r.(*x.Error)
		result = gin.H{
			"code":     e.Code,
			"message":  e.Message,
			"data":     e.Data,
			"trace_id": logger.GetId(),
		}
		break
	case x.Error:
		e := r.(x.Error)
		result = gin.H{
			"code":     e.Code,
			"message":  e.Message,
			"data":     e.Data,
			"trace_id": logger.GetId(),
		}
		break
	default:
		result = gin.H{
			"code":     x.UNKNOWN_ERRNO,
			"message":  "unknown error",
			"trace_id": logger.GetId(),
		}

		logger.LogFatal("unknown error", x.H{
			"errors": r,
		})
	}

	return
}
