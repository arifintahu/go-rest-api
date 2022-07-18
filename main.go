package main

import (
	"log"
	"os"

	"github.com/arifintahu/go-rest-api/app"
	"github.com/arifintahu/go-rest-api/app/config"
	"github.com/joho/godotenv"

	"github.com/unrolled/render"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	var (
		server = os.Getenv("SERVER")
		port = os.Getenv("PORT")
		appName = os.Getenv("APP_NAME")
	)

	logger := log.New(os.Stdout, appName, log.LstdFlags)
	render := render.New()

	appEnv := config.AppEnv{
		Logger: logger,
		Render: render,
		Server: server,
		Port: port,
		AppName: appName,
	}

	app.StartServer(appEnv)
}
