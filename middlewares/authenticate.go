package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/arifintahu/go-rest-api/dto"
	"github.com/arifintahu/go-rest-api/utils/jwt"
	"github.com/gin-gonic/gin"
)

var (
	ErrAuthorizationRequired = errors.New("Required Authorization")
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := strings.Replace(ctx.GetHeader("Authorization"), "Bearer ", "", -1)
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, dto.BaseErrorResponse(ErrAuthorizationRequired))
			ctx.Abort()
			return
		}

		userId, err := jwt.VerifyJWT(tokenString)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, dto.BaseErrorResponse(err))
			ctx.Abort()
			return
		}

		ctx.Set("userId", userId)
		ctx.Next()
	}
}
