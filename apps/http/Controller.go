package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"leaf/e"
	"leaf/mounts"
	
)

type Controller struct {
	ctx *gin.Context
	log x.Logger
}

func (c *Controller) Initialize(gtx *gin.Context, params interface{}) (ctx x.Context) {
	c.ctx = gtx
	c.init()

	if err := gtx.Bind(params); err != nil {
		x.ThrowError(e.ParamsFailed)
	}

	fmt.Printf("params: %T, %v\n", params, params)
	if err := mounts.Validator.Valid(params); err != nil {
		x.ThrowError(e.ParamsFailed, x.H{
			"errors": err,
		})
	}

	c.log.LogInfo("request_input", params)
	c.log.Params(params)
	return x.NewContextWithGin(gtx, c.log)
}
func (c *Controller) init() {
	log, _ := c.ctx.Get("log")
	c.log = log.(x.Logger)
}

func (c Controller) Response(code int, data interface{}, args ...interface{}) {
	if c.ctx.Writer.Written() {
		return
	}

	if r, ok := data.(x.Response); ok {
		// 越过中间件直接响应 不再在中间件进行额外操作
		c.ctx.AbortWithStatusJSON(code, r)
		c.log.LogInfo("request_output", r)
		return
	}

	c.ctx.AbortWithStatusJSON(code, data)
}
