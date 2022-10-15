package role

import (
	"github.com/arifintahu/go-rest-api/dto"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase IUseCase
}

type IController interface {
	CreateRole(ctx *gin.Context) (dto.BaseResponse, error)
	GetRoles(ctx *gin.Context) (dto.BaseResponse, error)
}

var _ IController = (*Controller)(nil)

func (c Controller) CreateRole(ctx *gin.Context) (dto.BaseResponse, error) {
	var body dto.RoleInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return dto.BaseResponse{}, err
	}

	data, err := c.useCase.CreateRole(&body)
	if err != nil {
		return dto.BaseResponse{}, err
	}

	res := dto.BaseResponse{
		Success: true,
		MessageTitle: "Success",
		Message: "Successfully create role!",
		Data: data,
	}
	return res, nil
}

func (c Controller) GetRoles(ctx *gin.Context) (dto.BaseResponse, error) {
	data, err := c.useCase.GetRoles()
	if err != nil {
		return dto.BaseResponse{}, err
	}

	res := dto.BaseResponse{
		Success: true,
		MessageTitle: "Success",
		Message: "Successfully get list roles!",
		Data: data,
	}
	return res, nil
}
