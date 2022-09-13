package app

import (
	"github.com/arifintahu/go-rest-api/app/config"
	"github.com/arifintahu/go-rest-api/app/routes"
	"github.com/arifintahu/go-rest-api/app/utils/handlers"
	"github.com/gin-gonic/gin"
)

func StartServer(appEnv config.AppEnv) {
	r := gin.Default()
	for _, route := range routes.RouteList {
		var handler gin.HandlerFunc
		handler = handlers.MakeHandler(appEnv, route.HandlerFunc)
		r.Handle(route.Method, route.Pattern, handler)
	}

	startupMessage := "===> Starting app "+ appEnv.AppName +""
	startupMessage = startupMessage + " on port " + appEnv.Port
	startupMessage = startupMessage + " in " + appEnv.Server + " mode."
	appEnv.Logger.Println(startupMessage)

	if appEnv.IsDevelopment {
		r.Run("localhost:" + appEnv.Port)
	} else {
		r.Run(":" + appEnv.Port)
	}
}
