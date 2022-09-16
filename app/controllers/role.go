package controllers

import (
	"net/http"

	"github.com/arifintahu/go-rest-api/app/config"
	"github.com/arifintahu/go-rest-api/app/dto"
	"github.com/arifintahu/go-rest-api/app/services"
	"github.com/arifintahu/go-rest-api/app/utils/handlers"
	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context, appEnv config.AppEnv) {
	var params dto.RoleInput
	if err := c.ShouldBindJSON(&params); err != nil {
		handlers.ResponseError(c, http.StatusUnprocessableEntity, err)
		return
	}

	s := services.NewRoleService(appEnv.DB)
	data, err := s.CreateRole(&params)
	if err != nil {
		handlers.ResponseError(c, http.StatusUnprocessableEntity, err)
		return
	}

	handlers.ResponseData(c, &gin.H{
		"message": "Successfully create role!",
		"data": data,
	})
}

func GetRoles(c *gin.Context, appEnv config.AppEnv) {
	s := services.NewRoleService(appEnv.DB)
	data, err := s.GetRoles()
	if err != nil {
		handlers.ResponseError(c, http.StatusUnprocessableEntity, err)
		return
	}

	handlers.ResponseData(c, &gin.H{
		"message": "Successfully get list roles!",
		"data": data,
	})
}
