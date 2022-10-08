package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	DbHost string
	DbUser string
	DbPassword string
	DbName string
	DbPort string
	DbTimeZone string
}

func NewDB(
	DbHost string,
	DbUser string,
	DbPassword string,
	DbName string,
	DbPort string,
	DbTimeZone string,
) *DBConfig {
	return &DBConfig{
		DbHost:DbHost,
		DbUser: DbUser,
		DbPassword: DbPassword,
		DbName: DbName,
		DbPort: DbPort,
		DbTimeZone: DbTimeZone,
	}
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
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
