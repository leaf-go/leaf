package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"leaf-go/apps/http"
	"leaf-go/params"
	"x"
)

type TestController struct {
	http.Controller
}

func (c TestController) Index(ctx *gin.Context) {
	param := &params.Test{}
	xc := c.Initialize(ctx, param)

	xc.LogInfo("hahaha", x.H{
		"error": "not found",
	})

	xc.LogError("hahaha", x.H{
		"error": "not found",
	})

	js, _ := json.Marshal(param)
	fmt.Println(js)

	c.Response(200, x.H{
		"success": true,
		"params":  param,
	})

}
