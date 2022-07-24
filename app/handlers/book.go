package handlers

import (
	"net/http"
	"strconv"

	"github.com/arifintahu/go-rest-api/app/config"

	"github.com/gorilla/mux"
)

func ListBooks(rw http.ResponseWriter, r *http.Request, appEnv config.AppEnv) {
	list, err := bookRepository.ListBooks()
	if err != nil {
		HandleError(rw, http.StatusNotFound, err)
	}
	appEnv.Render.JSON(rw, http.StatusOK, list)
}

func GetBook(rw http.ResponseWriter, r *http.Request, appEnv config.AppEnv) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	book, err := bookRepository.GetBook(id)
	if err != nil {
		HandleError(rw, http.StatusNotFound, err)
	}
	appEnv.Render.JSON(rw, http.StatusOK, book)
}
