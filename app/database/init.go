package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	DbHost string
	DbUser string
	DbPassword string
	DbName string
	DbPort string
	DbTimeZone string
}

func (dbConfig *DBConfig) InitConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		dbConfig.DbHost,
		dbConfig.DbUser,
		dbConfig.DbPassword,
		dbConfig.DbName,
		dbConfig.DbPort,
		dbConfig.DbTimeZone,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
