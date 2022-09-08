package routers

import (
	"example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func AlbumRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetAlbums)
}
