package main

import (
	"log"
	"os"

	"github.com/arifintahu/go-rest-api/entities"
	"github.com/arifintahu/go-rest-api/utils/db"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) error {
	err := db.Debug().AutoMigrate(
		&entities.Book{},
		&entities.Role{},
		&entities.User{},
	)
	if err != nil {
		return err
	}

	return nil
}

func main() {
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

	err = migrate(dbConn)
	if err != nil {
		log.Fatalf("Error migrate: %v", err)
	}

	sqlDB, err := dbConn.DB()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	sqlDB.Close()
}
