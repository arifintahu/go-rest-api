package config

import (
	"log"

	"github.com/unrolled/render"
	"gorm.io/gorm"
)

type AppEnv struct {
	Logger *log.Logger
	Render *render.Render
	Server string
	Port string
	AppName string
	DB *gorm.DB
}
