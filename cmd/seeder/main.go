package main

import (
	"log"
	"os"

	"github.com/arifintahu/go-rest-api/app/database"
	"github.com/arifintahu/go-rest-api/app/database/seed"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	dbConfig := database.DBConfig{
		DbHost:     os.Getenv("DB_HOST"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbPort:     os.Getenv("DB_PORT"),
		DbTimeZone: os.Getenv("DB_TIMEZONE"),
	}

	db, err := dbConfig.InitConnection()

	if err != nil {
		log.Fatal("Error initializing database connection")
	}

	err = seed.Load(db)
	if err != nil {
		log.Fatalf("Error seeder: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error seeder: %v", err)
	}

	sqlDB.Close()
}
