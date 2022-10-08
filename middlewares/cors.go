package middlewares

import "github.com/gin-gonic/gin"

func AllowCORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if origin := ctx.Request.Header.Get("Origin"); origin != "" {
			ctx.Header("Access-Control-Allow-Origin", origin)
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Menu-Slug, X-Origin-Path, X-Request-Id")
			ctx.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

			if ctx.Request.Method == "OPTIONS" {
				ctx.AbortWithStatus(204)
				return
			}
		}
		ctx.Next()
	}
}
