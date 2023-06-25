package routers

import (
	"github.com/gin-gonic/gin"
)

var sampleRouter SampleRouter

func Init() *gin.Engine {
	// v1 route group
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		// v1.Use(AuthMiddleware)
		sampleRouter.handle(v1.Group("/samples"))
	}

	return router
}
