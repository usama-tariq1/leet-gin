package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
}

// getAlbums responds with the list of all albums as JSON.
func (AlbumController AlbumController) GetAlbums(app *gin.Context) {

	// models.Lamp().Find(&albums) // method 1

	Lamp().Find(&albums) // method 2

	app.IndentedJSON(http.StatusOK, &albums)
}
