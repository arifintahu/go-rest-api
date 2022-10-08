package book

import (
	"strconv"

	"github.com/arifintahu/go-rest-api/dto"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase IUseCase
}

type IController interface {
	ListBooks(ctx *gin.Context) (dto.BaseResponse, error)
	GetBook(ctx *gin.Context) (dto.BaseResponse, error)
	AddBook(ctx *gin.Context) (dto.BaseResponse, error)
	UpdateBook(ctx *gin.Context) (dto.BaseResponse, error)
	DeleteBook(ctx *gin.Context) (dto.BaseResponse, error)
}

var _ IController = (*Controller)(nil)

func (c Controller) ListBooks(ctx *gin.Context) (dto.BaseResponse, error) {
	var res dto.BaseResponse
	data, err := c.useCase.ListBooks()
	if err != nil {
		return dto.BaseResponse{}, err
	}
	res = dto.BaseResponse{
		Success: true,
		MessageTitle: "Success",
		Message: "Successfully get list books!",
		Data: data,
	}
	return res, nil
}

func (c Controller) GetBook(ctx *gin.Context) (dto.BaseResponse, error) {
	var res dto.BaseResponse
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	data, err := c.useCase.GetBook(id)
	if err != nil {
		return dto.BaseResponse{}, err
	}
	res = dto.BaseResponse{
		Success: true,
		MessageTitle: "Success",
		Message: "Successfully get book!",
		Data: data,
	}
	return res, nil
}

func (c Controller) AddBook(ctx *gin.Context) (dto.BaseResponse, error) {
	var res dto.BaseResponse
	var body dto.BookInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return dto.BaseResponse{}, err
	}

	err := c.useCase.AddBook(&body)
	if err != nil {
		return dto.BaseResponse{}, err
	}

	res = dto.BaseResponse{
		Success: true,
		MessageTitle: "Success",
		Message: "Successfully add book!",
	}
	return res, nil
}

func (c Controller) UpdateBook(ctx *gin.Context) (dto.BaseResponse, error) {
	var res dto.BaseResponse
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var body dto.BookInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return dto.BaseResponse{}, err
	}

	err := c.useCase.UpdateBook(id, &body)

	if err != nil {
		return dto.BaseResponse{}, err
	}

	res = dto.BaseResponse{
		Success: true,
		MessageTitle: "Success",
		Message: "Successfully update book!",
	}
	return res, nil
}

func (c Controller) DeleteBook(ctx *gin.Context) (dto.BaseResponse, error) {
	var res dto.BaseResponse
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	
	err := c.useCase.DeleteBook(id)
	if err != nil {
		return dto.BaseResponse{}, err
	}
	res = dto.BaseResponse{
		Success: true,
		MessageTitle: "Success",
		Message: "Successfully delete book!",
	}
	return res, nil
}

