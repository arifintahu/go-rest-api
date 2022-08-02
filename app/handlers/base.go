package handlers

import (
	"io"
	"net/http"

	"github.com/arifintahu/go-rest-api/app/config"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, config.AppEnv)

func MakeHandler(appEnv config.AppEnv, fn func(http.ResponseWriter, *http.Request, config.AppEnv)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fn(rw, r, appEnv)
	}
}

func HandleError(rw http.ResponseWriter, status int, err error) {
	io.WriteString(rw, err.Error())
	rw.WriteHeader(status)
	return
}

func HealthCheck(rw http.ResponseWriter, r *http.Request, appEnv config.AppEnv) {
	io.WriteString(rw, "OK")
	rw.WriteHeader(http.StatusOK)
}
