package controllers

import (
	"net/http"

	"github.com/usama-tariq1/leet-gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(app *gin.Context) {
	var albums []models.Album
	DB = models.DB
	DB.Find(&albums)

	app.IndentedJSON(http.StatusOK, &albums)
}
