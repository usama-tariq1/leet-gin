package migrations

import (
	"github.com/usama-tariq1/leet-gin/helper"
	"github.com/usama-tariq1/leet-gin/models"
	"gorm.io/gorm"
)

// var DB *gorm.DB
var console helper.Console

var sample models.Sample

// list models here
func Migrate(DB *gorm.DB) {

	DB.AutoMigrate(sample)

}
