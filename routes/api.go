package routes

import (
	"github.com/gin-gonic/gin"
	"leaf-go/apps/http/api"
	mw "leaf-go/middleware"
	"x"
)

type API struct {
}

func (r *API) init(app x.IApplication) *gin.Engine {
	handler := app.Handler()
	router := handler.(*gin.Engine)
	router.NoRoute(mw.ErrorNotFound())
	router.NoMethod(mw.ErrorNotFound())
	router.Use(mw.Init(), mw.Cors(), mw.Catch())

	return router
}

func (r API) Boot(app x.IApplication) {
	router := r.init(app)

	normal := router.Group("api")
	{
		c := &api.TestController{}
		normal.POST("/index", c.Index)
	}

	router.GET("/xxx", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"xxx": "123",
		})
	})
}
