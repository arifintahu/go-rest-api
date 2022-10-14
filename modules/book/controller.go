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
	CreateBook(ctx *gin.Context) (dto.BaseResponse, error)
	GetBooks(ctx *gin.Context) (dto.BaseResponse, error)
	GetBookDetail(ctx *gin.Context) (dto.BaseResponse, error)
	UpdateBook(ctx *gin.Context) (dto.BaseResponse, error)
	DeleteBook(ctx *gin.Context) (dto.BaseResponse, error)
}

var _ IController = (*Controller)(nil)

func (c Controller) CreateBook(ctx *gin.Context) (dto.BaseResponse, error) {
	var res dto.BaseResponse
	var body dto.BookInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return dto.BaseResponse{}, err
	}

	data, err := c.useCase.CreateBook(&body)
	if err != nil {
		return dto.BaseResponse{}, err
	}

	res = dto.BaseResponse{
		Success: true,
		MessageTitle: "Success",
		Message: "Successfully create book!",
		Data: data,
	}
	return res, nil
}

func (c Controller) GetBooks(ctx *gin.Context) (dto.BaseResponse, error) {
	var res dto.BaseResponse
	data, err := c.useCase.GetBooks()
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

func (c Controller) GetBookDetail(ctx *gin.Context) (dto.BaseResponse, error) {
	var res dto.BaseResponse
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	data, err := c.useCase.GetBookDetail(id)
	if err != nil {
		return dto.BaseResponse{}, err
	}
	res = dto.BaseResponse{
		Success: true,
		MessageTitle: "Success",
		Message: "Successfully get book detail!",
		Data: data,
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

	data, err := c.useCase.UpdateBook(id, &body)

	if err != nil {
		return dto.BaseResponse{}, err
	}

	res = dto.BaseResponse{
		Success: true,
		MessageTitle: "Success",
		Message: "Successfully update book!",
		Data: data,
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

