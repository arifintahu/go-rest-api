package controllers

import (
	"net/http"
	"strconv"

	"github.com/arifintahu/go-rest-api/app/config"
	"github.com/arifintahu/go-rest-api/app/dto"
	"github.com/arifintahu/go-rest-api/app/services"
	"github.com/arifintahu/go-rest-api/app/utils/handlers"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context, appEnv config.AppEnv) {
	c.JSON(http.StatusOK, gin.H{})
}

func ListBooks(c *gin.Context, appEnv config.AppEnv) {
	s := services.NewBookService(appEnv.DB)
	data, err := s.ListBooks()
	if err != nil {
		handlers.ResponseError(c, http.StatusUnprocessableEntity, err)
		return
	}
	handlers.ResponseData(c, &gin.H{
		"message": "Successfully get list books!",
		"data": data,
	})
}

func GetBook(c *gin.Context, appEnv config.AppEnv) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	s := services.NewBookService(appEnv.DB)
	data, err := s.GetBook(id)
	if err != nil {
		handlers.ResponseError(c, http.StatusUnprocessableEntity, err)
		return
	}
	handlers.ResponseData(c, &gin.H{
		"message": "Successfully get book!",
		"data": data,
	})
}

func AddBook(c *gin.Context, appEnv config.AppEnv) {
	var body dto.BookInput
	if err := c.ShouldBindJSON(&body); err != nil {
		handlers.ResponseError(c, http.StatusUnprocessableEntity, err)
		return
	}

	s := services.NewBookService(appEnv.DB)
	err := s.AddBook(&body)
	if err != nil {
		handlers.ResponseError(c, http.StatusUnprocessableEntity, err)
		return
	}

	handlers.ResponseData(c, &gin.H{
		"message": "Successfully add book!",
	})
}

func UpdateBook(c *gin.Context, appEnv config.AppEnv) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var body dto.BookInput
	if err := c.ShouldBindJSON(&body); err != nil {
		handlers.ResponseError(c, http.StatusUnprocessableEntity, err)
		return
	}

	s := services.NewBookService(appEnv.DB)
	err := s.UpdateBook(id, &body)

	if err != nil {
		handlers.ResponseError(c, http.StatusUnprocessableEntity, err)
		return
	}

	handlers.ResponseData(c, &gin.H{
		"message": "Successfully update book!",
	})
}

func DeleteBook(c *gin.Context, appEnv config.AppEnv) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	
	s := services.NewBookService(appEnv.DB)
	err := s.DeleteBook(id)
	if err != nil {
		handlers.ResponseError(c, http.StatusUnprocessableEntity, err)
		return
	}
	handlers.ResponseData(c, &gin.H{"message": "Successfully delete book!"})
}
