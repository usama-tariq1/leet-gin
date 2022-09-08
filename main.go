package main

import (
	"example/web-service-gin/config"
	"example/web-service-gin/models"
	"example/web-service-gin/routers"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// app start
	app := routers.Init()

	// get env
	env := config.Env{}
	port := env.Get("PORT")

	// connect Database
	models.ConnectDatabase()

	app.Run(fmt.Sprintf("localhost:%s", port))
}
