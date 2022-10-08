package main

import (
	"log"
	"os"

	"github.com/arifintahu/go-rest-api/app"

	"github.com/arifintahu/go-rest-api/utils/db"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var (
		env = os.Getenv("ENV")
		port = os.Getenv("PORT")
		appName = os.Getenv("APP_NAME")
	)

	logger := log.New(os.Stdout, appName, log.LstdFlags)

	gormDB := db.NewDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_TIMEZONE"),
	)

	dbConn, err := gormDB.InitConnection()

	if err != nil {
		log.Fatal("Error initializing database connection")
	}

	application := app.NewApp(
		logger,
		appName,
		dbConn,
		env,
		port,
	)

	application.StartServer()
}
