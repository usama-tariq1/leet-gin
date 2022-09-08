package main

import (
	"fmt"
	"log"

	"github.com/usama-tariq1/leet-gin/config"
	"github.com/usama-tariq1/leet-gin/models"
	"github.com/usama-tariq1/leet-gin/routers"

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
