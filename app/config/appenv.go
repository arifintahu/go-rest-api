package config

import (
	"log"

	"gorm.io/gorm"
)

type AppEnv struct {
	Logger *log.Logger
	Server string
	Port string
	AppName string
	IsDevelopment bool
	DB *gorm.DB
}
