package main

import (
	"log"
	"os"

	"github.com/arifintahu/go-rest-api/app/models"
	"github.com/arifintahu/go-rest-api/app/utils/database"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) error {
	err := db.Debug().AutoMigrate(&models.Book{})
	if err != nil {
		return err
	}

	return nil
}

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

	err = migrate(db)
	if err != nil {
		log.Fatalf("Error migrate: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	sqlDB.Close()
}
