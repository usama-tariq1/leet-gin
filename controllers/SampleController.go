package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/usama-tariq1/leet-gin/helper"
	"github.com/usama-tariq1/leet-gin/models"
)

type SampleController struct {
}

var ()

func (SampleController) Index(app *gin.Context) {
	page := app.Query("page")
	limit := app.Query("limit")

	// Set default values if not provided
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}

	// Convert page and limit to integers
	pageInt := helper.ConvertToInt(page)
	limitInt := helper.ConvertToInt(limit)

	// Retrieve the list from the model with pagination
	var Sample models.Sample
	list, totalCount, err := Sample.GetList(pageInt, limitInt)
	if err != nil {
		app.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve Sample list",
		})
		return
	}

	// Prepare the response
	response := gin.H{
		"list":       list,
		"page":       pageInt,
		"limit":      limitInt,
		"totalCount": totalCount,
	}

	app.JSON(http.StatusOK, response)
}

func (SampleController) Create(app *gin.Context) {
	// Handle create operation using GORM and Sample model
	var Sample models.Sample
	app.BindJSON(&Sample)
	err := Sample.Create(&Sample)
	if err != nil {
		app.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create record",
		})
		return
	}

	app.IndentedJSON(http.StatusOK, Sample)
}

func (SampleController) Read(app *gin.Context) {
	// Handle read operation using GORM and Sample model
	var Sample models.Sample

	id := uuid.MustParse(app.Param("id"))
	result, err := Sample.GetByID(id)
	if err != nil {
		app.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read record",
		})
		return
	}

	app.IndentedJSON(http.StatusOK, result)
}

func (SampleController) Update(app *gin.Context) {
	// Handle update operation using GORM and Sample model
	var Sample models.Sample

	id := uuid.MustParse(app.Param("id"))

	app.BindJSON(&Sample)

	err := Sample.Update(&Sample, id)
	if err != nil {
		app.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update record",
		})
		return
	}

	result, err := Sample.GetByID(id)
	if err != nil {
		app.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read record",
		})
		return
	}

	app.IndentedJSON(http.StatusOK, result)
}

func (SampleController) Delete(app *gin.Context) {
	// Handle delete operation using GORM and Sample model
	var Sample models.Sample

	id := uuid.MustParse(app.Param("id"))
	err := Sample.Delete(id)
	if err != nil {
		if err.Error() == "record not found" {
			app.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Failed to delete, record not found",
			})
		} else {
			app.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to delete record",
			})
		}
		return
	}

	app.IndentedJSON(http.StatusOK, "Delete Successfully")
}
