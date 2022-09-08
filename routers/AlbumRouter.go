package routers

import (
	"github.com/usama-tariq1/leet-gin/controllers"

	"github.com/gin-gonic/gin"
)

func AlbumRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetAlbums)
}
