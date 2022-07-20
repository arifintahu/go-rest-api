package handlers

import (
	"io"
	"net/http"

	"github.com/arifintahu/go-rest-api/app/config"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, config.AppEnv)

func MakeHandler(appEnv config.AppEnv, fn func(http.ResponseWriter, *http.Request, config.AppEnv)) http.HandlerFunc {
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
