package handlers

import (
	"io"
	"log"
	"net/http"

	"github.com/unrolled/render"
)

type Handler struct {
	Logger *log.Logger
	Render *render.Render
}

func InitHandler(logger *log.Logger, render *render.Render) *Handler {
	return &Handler{
		Logger: logger,
		Render: render,
	}
}

func (handler *Handler) PingConnection(rw http.ResponseWriter, r *http.Request) {
	handler.Logger.Println("Handler: pingConnection")
	io.WriteString(rw, "OK")
	rw.WriteHeader(http.StatusOK)
}
