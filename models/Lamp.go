package models

import "gorm.io/gorm"

func Lamp() *gorm.DB {
	return DB
}
