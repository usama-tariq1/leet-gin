package routers

import (
	"github.com/gin-gonic/gin"
)

var albumRouter AlbumRouter

func Init() *gin.Engine {
	// v1 route group
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		// v1.Use(AuthMiddleware)
		albumRouter.handle(v1.Group("/albums"))
	}

	return router
}
