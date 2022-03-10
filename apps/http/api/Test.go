package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leaf-go/x"
	"leaf/apps/http"
	"leaf/params"
	
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
