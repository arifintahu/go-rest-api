package app

import (
	"log"

	"github.com/arifintahu/go-rest-api/middlewares"
	"github.com/arifintahu/go-rest-api/modules/book"
	"github.com/arifintahu/go-rest-api/modules/role"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	logger *log.Logger
	name string // application name
	db *gorm.DB // Postgres DB
	env string // application environment {local, staging, production}
	port string //application running port
}

func NewApp(
	logger *log.Logger,
	name string,
	db *gorm.DB,
	env string,
	port string,
) *App {
	return &App{
		logger: logger,
		name: name,
		db: db,
		env: env,
		port: port,
	}
}

func (app *App) StartServer() {
	r := gin.Default()
	r.Use(
		middlewares.AllowCORS(),
	)

	roleHandler := role.NewRequestHandler(app.db, app.logger)
	roleHandler.Handle(r)

	bookHandler := book.NewRequestHandler(app.db, app.logger)
	bookHandler.Handle(r)
	
	startupMessage := "===> Starting app "+ app.name +""
	startupMessage = startupMessage + " on port " + app.port
	startupMessage = startupMessage + " in " + app.env + " mode."
	app.logger.Println(startupMessage)

	err := r.Run(":" + app.port)

	if err != nil {
		app.logger.Panicln(err)
		return
	}
}
