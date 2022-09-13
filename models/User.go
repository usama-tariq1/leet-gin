package models

import "gorm.io/gorm"

type User struct {
	// gorm.Model include ID , CreatedAt , UpdatedAt , DeletedAt
	gorm.Model

	// list you table columns here
	// Name string `json:"name"`
}
