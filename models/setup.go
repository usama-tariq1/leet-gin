package models

import (
	"example/web-service-gin/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	env := config.Env{}
	host := env.Get("DB_HOST")
	user := env.Get("DB_USER")
	pass := env.Get("DB_PASS")
	db_name := env.Get("DB_Name")
	port := env.Get("DB_PORT")
	// sslmode := env.Get("DB_SSL")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, pass, db_name, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Connected to Database")

	fmt.Println("Starting Migration")
	db.AutoMigrate(&Album{})
	fmt.Println("Migration Completed")

	DB = db
}
