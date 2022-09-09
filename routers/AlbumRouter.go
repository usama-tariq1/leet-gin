package routers

import (
	"github.com/usama-tariq1/leet-gin/controllers"

	"github.com/gin-gonic/gin"
)

type AlbumRouter struct {
}

var AlbumController controllers.AlbumController

func (AlbumRouter AlbumRouter) handle(router *gin.RouterGroup) {
	router.GET("/", AlbumController.GetAlbums)
}
