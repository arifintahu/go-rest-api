package main

import (
	"log"
	"os"

	"github.com/arifintahu/go-rest-api/app"
	"github.com/arifintahu/go-rest-api/app/config"

	"github.com/unrolled/render"
)

func main() {
	var (
		env = "LOCAL"
		port = "3000"
		appName = "rest-api"
	)
	logger := log.New(os.Stdout, appName, log.LstdFlags)
	render := render.New()

	appEnv := config.AppEnv{
		Logger: logger,
		Render: render,
		Env: env,
		Port: port,
		AppName: appName,
	}

	app.StartServer(appEnv)
}
