package handlers

import (
	"net/http"

	"github.com/arifintahu/go-rest-api/app/config"
	"github.com/gin-gonic/gin"
)

type HandlerFunc func(*gin.Context, config.AppEnv)

func MakeHandler(appEnv config.AppEnv, fn func(*gin.Context, config.AppEnv)) gin.HandlerFunc {
	return func(c *gin.Context) {
		fn(c, appEnv)
	}
}

func ResponseError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"error": true,
		"message" : err.Error(),
	})
}

func ResponseData(c *gin.Context, data *gin.H ) {
	c.JSON(http.StatusOK, data)
}
