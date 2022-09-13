package controllers

import (
	"net/http"
	"strconv"

	"github.com/arifintahu/go-rest-api/app/config"
	"github.com/arifintahu/go-rest-api/app/models"
	"github.com/arifintahu/go-rest-api/app/repositories"
	"github.com/arifintahu/go-rest-api/app/utils/handlers"
	"github.com/gin-gonic/gin"
)

type BookInput struct {
	Title     	string 	`json:"title"`
	Author    	string 	`json:"author"`
	Page      	uint16 	`json:"page"`
	Publisher 	string 	`json:"publisher"`
	Quantity  	uint16 	`json:"quantity"`
}

func HealthCheck(c *gin.Context, appEnv config.AppEnv) {
	c.JSON(http.StatusOK, gin.H{})
}

func ListBooks(c *gin.Context, appEnv config.AppEnv) {
	bookRepository := repositories.NewBookRepository(appEnv.DB)
	list, err := bookRepository.ListBooks()
	if err != nil {
		handlers.ResponseError(c, http.StatusNotFound, err)
		return
	}
	handlers.ResponseData(c, &gin.H{
		"message": "Successfully get list books!",
		"data": list,
	})
}

func GetBook(c *gin.Context, appEnv config.AppEnv) {
	bookRepository := repositories.NewBookRepository(appEnv.DB)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	book, err := bookRepository.GetBook(id)
	if err != nil {
		handlers.ResponseError(c, http.StatusNotFound, err)
		return
	}
	handlers.ResponseData(c, &gin.H{
		"message": "Successfully get book!",
		"data": book,
	})
}

func AddBook(c *gin.Context, appEnv config.AppEnv) {
	bookRepository := repositories.NewBookRepository(appEnv.DB)
	var input BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		handlers.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	book := models.Book{
		Title: input.Title,
		Author: input.Author,
		Page: input.Page,
		Publisher: input.Publisher,
		Quantity: input.Quantity,
	}

	err := bookRepository.AddBook(&book)
	if err != nil {
		handlers.ResponseError(c, http.StatusUnprocessableEntity, err)
		return
	}

	handlers.ResponseData(c, &gin.H{
		"message": "Successfully add book!",
		"data": book,
	})
}

func UpdateBook(c *gin.Context, appEnv config.AppEnv) {
	bookRepository := repositories.NewBookRepository(appEnv.DB)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	_, err := bookRepository.GetBook(id)
	if err != nil {
		handlers.ResponseError(c, http.StatusNotFound, err)
		return
	}

	var input BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		handlers.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	book := models.Book{
		Title: input.Title,
		Author: input.Author,
		Page: input.Page,
		Publisher: input.Publisher,
		Quantity: input.Quantity,
	}

	err = bookRepository.UpdateBook(id, &book)
	if err != nil {
		handlers.ResponseError(c, http.StatusUnprocessableEntity, err)
		return
	}

	handlers.ResponseData(c, &gin.H{
		"message": "Successfully update book!",
		"data": book,
	})
}

func DeleteBook(c *gin.Context, appEnv config.AppEnv) {
	bookRepository := repositories.NewBookRepository(appEnv.DB)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	_, err := bookRepository.GetBook(id)
	if err != nil {
		handlers.ResponseError(c, http.StatusNotFound, err)
		return
	}

	err = bookRepository.DeleteBook(id)
	if err != nil {
		handlers.ResponseError(c, http.StatusUnprocessableEntity, err)
		return
	}
	handlers.ResponseData(c, &gin.H{"message": "Successfully delete book!"})
}
