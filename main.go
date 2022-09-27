package main

import ( // "fmt"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/usama-tariq1/leet-gin/config"
	"github.com/usama-tariq1/leet-gin/helper"
	leetgin "github.com/usama-tariq1/leet-gin/leet-gin"
	"github.com/usama-tariq1/leet-gin/migrations"
	"github.com/usama-tariq1/leet-gin/models"
	"github.com/usama-tariq1/leet-gin/routers"
)

var console = helper.Console{}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// app start
	leetgin.Welcome()
	app := routers.Init()

	// get env
	env := config.Env{}
	port := env.Get("PORT")

	// connect Database
	connection := models.ConnectDatabase()
	console.Log("Debug", "** Starting Migration **")
	migrations.Migrate(connection)
	console.Log("Debug", "** Migration Completed **")

	console.Log("Debug", fmt.Sprintf("--> Server starting on localhost:%s", port))

	app.Run(fmt.Sprintf("localhost:%s", port))

}
