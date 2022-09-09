package controllers

import (
	"net/http"

	"github.com/usama-tariq1/leet-gin/models"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
}

var (
	DB     = models.DB
	albums []models.Album
	album  models.Album
)

// getAlbums responds with the list of all albums as JSON.
func (AlbumController AlbumController) GetAlbums(app *gin.Context) {

	models.Lamp().Find(&albums) // method 1

	// DB.Find(&album) // method 2

	app.IndentedJSON(http.StatusOK, &albums)
}
