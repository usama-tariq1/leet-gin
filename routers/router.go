package routers

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine{
	// Simple group: v1
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		// v1.Use(AuthMiddleware)
		AlbumRoutes(v1.Group("/albums"))
	}

	return router
}