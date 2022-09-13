package controllers

import (
	"github.com/usama-tariq1/leet-gin/models"
	"gorm.io/gorm"
)

var (
	albums []models.Album
	album  models.Album
)

type Controller struct {
}

func Lamp() *gorm.DB {
	return models.DB
}
