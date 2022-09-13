package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

var ()

func (UserController UserController) Index(app *gin.Context) {

	app.IndentedJSON(http.StatusOK, "Index Called")
}

func (UserController UserController) Create(app *gin.Context) {

	app.IndentedJSON(http.StatusOK, "Create Called")

}

func (UserController UserController) Read(app *gin.Context) {

	app.IndentedJSON(http.StatusOK, "Read Called")

}

func (UserController UserController) Update(app *gin.Context) {

	app.IndentedJSON(http.StatusOK, "Update Called")

}

func (UserController UserController) Delete(app *gin.Context) {

	app.IndentedJSON(http.StatusOK, "Delete Called")

}
