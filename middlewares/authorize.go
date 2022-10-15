package middlewares

import (
	"errors"
	"net/http"

	"github.com/arifintahu/go-rest-api/dto"
	"github.com/arifintahu/go-rest-api/utils/jwt"
	"github.com/gin-gonic/gin"
)

var (
	ErrAuthorizationForbidden = errors.New("Forbidden Authorization")
)

type AuthData struct {
	jwt.Claims
}

func Authorize(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authDataRaw, exists := ctx.Get("authData")
		authData, isMap := authDataRaw.(map[string]interface{})
		if !exists || !isMap {
			ctx.JSON(http.StatusForbidden, dto.BaseErrorResponse(ErrAuthorizationForbidden))
			ctx.Abort()
			return
		}
		
		isAuthorized := containRole(roles, authData["role_slug"].(string))
		if !isAuthorized {
			ctx.JSON(http.StatusForbidden, dto.BaseErrorResponse(ErrAuthorizationForbidden))
			ctx.Abort()
			return
		} 
		
		ctx.Next()
	}
}

func containRole(roles []string, role string) bool {
	for i := range roles {
		if roles[i] == role {
			return true
		}
	}

	return false
}
