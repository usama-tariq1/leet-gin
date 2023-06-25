package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/usama-tariq1/leet-gin/controllers"
)

type SampleRouter struct {
}

var SampleController controllers.SampleController

func (SampleRouter SampleRouter) handle(router *gin.RouterGroup) {
	router.GET("", SampleController.Index)
	router.POST("", SampleController.Create)
	router.GET("/:id", SampleController.Read)
	router.PATCH("/:id", SampleController.Update)
	router.DELETE("/:id", SampleController.Delete)
}
