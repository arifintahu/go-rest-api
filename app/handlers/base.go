package handlers

import (
	"io"
	"net/http"

	"github.com/arifintahu/go-rest-api/app/config"
	"github.com/arifintahu/go-rest-api/app/repositories"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, config.AppEnv)

var bookRepository *repositories.BookRepository

func MakeHandler(appEnv config.AppEnv, fn func(http.ResponseWriter, *http.Request, config.AppEnv)) http.HandlerFunc {
	bookRepository = &repositories.BookRepository{
		DB: appEnv.DB,
	}
	
	return func(rw http.ResponseWriter, r *http.Request) {
		// return function with AppEnv
		fn(rw, r, appEnv)
	}
}

func PingConnection(rw http.ResponseWriter, r *http.Request, appEnv config.AppEnv) {
	appEnv.Logger.Println("Handler: pingConnection")
	io.WriteString(rw, "OK")
	rw.WriteHeader(http.StatusOK)
}
