package user

import (
	"strconv"

	"github.com/arifintahu/go-rest-api/dto"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase IUseCase
}

type IController interface {
	CreateUser(ctx *gin.Context) (dto.BaseResponse, error)
	GetUsers(ctx *gin.Context) (dto.BaseResponse, error)
	GetUserDetail(ctx *gin.Context) (dto.BaseResponse, error)
	UpdateUser(ctx *gin.Context) (dto.BaseResponse, error)
	DeleteUser(ctx *gin.Context) (dto.BaseResponse, error)
}

var _ IController = (*Controller)(nil)

func (c Controller) CreateUser(ctx *gin.Context) (dto.BaseResponse, error) {
	var body dto.UserInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return dto.BaseResponse{}, err
	}

	data, err := c.useCase.CreateUser(&body)
	if err != nil {
		return dto.BaseResponse{}, err
	}

	res := dto.BaseResponse{
		Success:      true,
		MessageTitle: "Success",
		Message:      "Successfully create user!",
		Data:         data,
	}
	return res, nil
}

func (c Controller) GetUsers(ctx *gin.Context) (dto.BaseResponse, error) {
	var query dto.UserListQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		return dto.BaseResponse{}, err
	}

	data, total, err := c.useCase.GetUsers(&query)
	if err != nil {
		return dto.BaseResponse{}, err
	}

	res := dto.BaseResponse{
		Success: true,
		MessageTitle: "Success",
		Message: "Successfully get list users!",
		Data: data,
		Total: total,
	}
	return res, nil
}

func (c Controller) GetUserDetail(ctx *gin.Context) (dto.BaseResponse, error) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)

	data, err := c.useCase.GetUserDetail(id)
	if err != nil {
		return dto.BaseResponse{}, err
	}
	res := dto.BaseResponse{
		Success: true,
		MessageTitle: "Success",
		Message: "Successfully get user detail!",
		Data: data,
	}
	return res, nil
}

func (c Controller) UpdateUser(ctx *gin.Context) (dto.BaseResponse, error) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	var body dto.UserUpdate
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return dto.BaseResponse{}, err
	}

	data, err := c.useCase.UpdateUser(id, &body)

	if err != nil {
		return dto.BaseResponse{}, err
	}

	res := dto.BaseResponse{
		Success: true,
		MessageTitle: "Success",
		Message: "Successfully update user!",
		Data: data,
	}
	return res, nil
}

func (c Controller) DeleteUser(ctx *gin.Context) (dto.BaseResponse, error) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	
	err := c.useCase.DeleteUser(id)
	if err != nil {
		return dto.BaseResponse{}, err
	}
	res := dto.BaseResponse{
		Success: true,
		MessageTitle: "Success",
		Message: "Successfully delete user!",
	}
	return res, nil
}
