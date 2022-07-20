package main

import (
	"log"
	"os"

	"github.com/arifintahu/go-rest-api/app"
	"github.com/arifintahu/go-rest-api/app/config"
	"github.com/joho/godotenv"

	"github.com/arifintahu/go-rest-api/app/database"
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

	dbConfig := database.DBConfig{
		DbHost: os.Getenv("DB_HOST"),
		DbUser: os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName: os.Getenv("DB_NAME"),
		DbPort: os.Getenv("DB_PORT"),
		DbTimeZone: os.Getenv("DB_TIMEZONE"),
	}

	db, err := dbConfig.InitConnection()

	if err != nil {
		log.Fatal("Error initializing database connection")
	}

	appEnv := config.AppEnv{
		Logger: logger,
		Render: render,
		Server: server,
		Port: port,
		AppName: appName,
		DB: db,
	}

	app.StartServer(appEnv)
}
