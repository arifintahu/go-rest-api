package config

import (
	"log"

	"github.com/unrolled/render"
)

type AppEnv struct {
	Logger *log.Logger
	Render *render.Render
	Server string
	Port string
	AppName string
}
