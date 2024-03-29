package account

import (
	"github.com/arifintahu/go-rest-api/dto"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase IUseCase
}

type IController interface {
	Login(ctx *gin.Context) (dto.BaseResponse, error)
}


var _ IController = (*Controller)(nil)

func (c Controller) Login(ctx *gin.Context) (dto.BaseResponse, error) {
	var body dto.AccountLogin
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return dto.BaseResponse{}, err
	}

	data, err := c.useCase.Login(&body)
	if err != nil {
		return dto.BaseResponse{}, err
	}

	res := dto.BaseResponse{
		Success:      true,
		MessageTitle: "Success",
		Message:      "Successfully login!",
		Data:         data,
	}
	return res, nil
}
