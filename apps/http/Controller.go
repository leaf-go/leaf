package http

import (
	"github.com/gin-gonic/gin"
	"x"
)

type Controller struct{}

func (c *Controller) Initialize(context *gin.Context, params interface{}) (ctx x.Context) {

	return  x.NewContextWithGin(context)
}
