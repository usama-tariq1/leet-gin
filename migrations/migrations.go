package migrations

import (
	leetGin "github.com/usama-tariq1/leet-gin/helper"
	"github.com/usama-tariq1/leet-gin/models"
	"gorm.io/gorm"
)

var console leetGin.Console

var sample models.Sample

// list models here
func Migrate(DB *gorm.DB) {

	DB.AutoMigrate(sample)

}
