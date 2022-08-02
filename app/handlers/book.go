package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/arifintahu/go-rest-api/app/config"
	"github.com/arifintahu/go-rest-api/app/models"
	"github.com/arifintahu/go-rest-api/app/repositories"

	"github.com/gorilla/mux"
)

func ListBooks(rw http.ResponseWriter, r *http.Request, appEnv config.AppEnv) {
	bookRepository := repositories.NewBookRepository(appEnv.DB)
	list, err := bookRepository.ListBooks()
	if err != nil {
		HandleError(rw, http.StatusNotFound, err)
	}
	appEnv.Render.JSON(rw, http.StatusOK, list)
}

func GetBook(rw http.ResponseWriter, r *http.Request, appEnv config.AppEnv) {
	bookRepository := repositories.NewBookRepository(appEnv.DB)
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	book, err := bookRepository.GetBook(id)
	if err != nil {
		HandleError(rw, http.StatusNotFound, err)
	}
	appEnv.Render.JSON(rw, http.StatusOK, book)
}

func AddBook(rw http.ResponseWriter, r *http.Request, appEnv config.AppEnv) {
	bookRepository := repositories.NewBookRepository(appEnv.DB)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HandleError(rw, http.StatusUnprocessableEntity, err)
	}

	book := models.Book{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		HandleError(rw, http.StatusUnprocessableEntity, err)
	}

	err = bookRepository.AddBook(&book)
	if err != nil {
		HandleError(rw, http.StatusUnprocessableEntity, err)
	}

	appEnv.Render.JSON(rw, http.StatusOK, book)
}

func UpdateBook(rw http.ResponseWriter, r *http.Request, appEnv config.AppEnv) {
	bookRepository := repositories.NewBookRepository(appEnv.DB)
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	_, err := bookRepository.GetBook(id)
	if err != nil {
		HandleError(rw, http.StatusNotFound, err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HandleError(rw, http.StatusUnprocessableEntity, err)
	}

	book := models.Book{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		HandleError(rw, http.StatusUnprocessableEntity, err)
	}

	err = bookRepository.UpdateBook(id, &book)
	if err != nil {
		HandleError(rw, http.StatusUnprocessableEntity, err)
	}

	appEnv.Render.JSON(rw, http.StatusOK, book)
}

func DeleteBook(rw http.ResponseWriter, r *http.Request, appEnv config.AppEnv) {
	bookRepository := repositories.NewBookRepository(appEnv.DB)
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)

	_, err := bookRepository.GetBook(id)
	if err != nil {
		HandleError(rw, http.StatusNotFound, err)
	}

	err = bookRepository.DeleteBook(id)
	if err != nil {
		HandleError(rw, http.StatusNotFound, err)
	}
	appEnv.Render.JSON(rw, http.StatusOK, map[string]interface{}{})
}
