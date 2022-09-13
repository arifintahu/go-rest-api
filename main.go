package main

import (
	"log"
	"os"

	"github.com/arifintahu/go-rest-api/app"
	"github.com/arifintahu/go-rest-api/app/config"

	"github.com/arifintahu/go-rest-api/app/utils/database"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var (
		server = os.Getenv("SERVER")
		port = os.Getenv("PORT")
		appName = os.Getenv("APP_NAME")
	)

	logger := log.New(os.Stdout, appName, log.LstdFlags)

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
		Server: server,
		Port: port,
		AppName: appName,
		DB: db,
		IsDevelopment: server == "local",
	}

	app.StartServer(appEnv)
}
