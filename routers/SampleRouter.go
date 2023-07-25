package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/usama-tariq1/leet-gin/controllers"
	// "github.com/usama-tariq1/leet-gin/middlewares"
)

type SampleRouter struct {
}

func (SampleRouter SampleRouter) handle(router *gin.RouterGroup) {
	var SampleController controllers.SampleController
	// router.Use(middlewares.Sample())

	router.GET("", SampleController.Index)
	router.POST("", SampleController.Create)
	router.GET("/:id", SampleController.Read)
	router.PATCH("/:id", SampleController.Update)
	router.DELETE("/:id", SampleController.Delete)
}
