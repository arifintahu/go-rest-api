package handlers

import (
	"go-rest-api-2/services"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var bookService *services.BookService

func (handler *Handler) ListBooks(rw http.ResponseWriter, r *http.Request) {
	handler.Logger.Println("Handler: ListBooks")
	list, err := bookService.ListBooks()
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		io.WriteString(rw, err.Error())
		rw.WriteHeader(http.StatusNotFound)
	}
	handler.Render.JSON(rw, http.StatusOK, list)
}

func (handler *Handler) GetBook(rw http.ResponseWriter, r *http.Request) {
	handler.Logger.Println("Handler: GetBook")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	book, err := bookService.GetBook(id)
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		io.WriteString(rw, err.Error())
		rw.WriteHeader(http.StatusNotFound)
	}
	handler.Render.JSON(rw, http.StatusOK, book)
}
