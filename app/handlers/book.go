package handlers

import (
	"io"
	"net/http"
	"strconv"

	"github.com/arifintahu/go-rest-api/app/config"
	"github.com/arifintahu/go-rest-api/app/services"

	"github.com/gorilla/mux"
)

var bookService *services.BookService

func ListBooks(rw http.ResponseWriter, r *http.Request, appEnv config.AppEnv) {
	list, err := bookService.ListBooks()
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		io.WriteString(rw, err.Error())
		rw.WriteHeader(http.StatusNotFound)
	}
	appEnv.Render.JSON(rw, http.StatusOK, list)
}

func GetBook(rw http.ResponseWriter, r *http.Request, appEnv config.AppEnv) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	book, err := bookService.GetBook(id)
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		io.WriteString(rw, err.Error())
		rw.WriteHeader(http.StatusNotFound)
	}
	appEnv.Render.JSON(rw, http.StatusOK, book)
}
