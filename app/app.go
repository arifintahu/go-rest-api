package app

import (
	"net/http"

	"github.com/arifintahu/go-rest-api/app/config"
	"github.com/arifintahu/go-rest-api/app/handlers"
	"github.com/arifintahu/go-rest-api/app/routes"
	"github.com/gorilla/mux"
	"github.com/unrolled/secure"
	"github.com/urfave/negroni"
)

const LOCAL string = "LOCAL"

func StartServer(appEnv config.AppEnv) {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes.RouteList {
		var handler http.Handler
		handler = handlers.MakeHandler(appEnv, route.HandlerFunc)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	var isDevelopment = false

	if appEnv.Server == LOCAL {
		isDevelopment = true
	}

	secureMiddleware := secure.New(secure.Options{
		// This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains
		// options to be ignored during development. When deploying to production,
		// be sure to set this to false.
		IsDevelopment: isDevelopment,
		// AllowedHosts is a list of fully qualified domain names that are allowed (CORS)
		AllowedHosts: []string{},
		// If ContentTypeNosniff is true, adds the X-Content-Type-Options header
		// with the value `nosniff`. Default is false.
		ContentTypeNosniff: true,
		// If BrowserXssFilter is true, adds the X-XSS-Protection header with the
		// value `1; mode=block`. Default is false.
		BrowserXssFilter: true,
	})

	// Start server
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.HandlerFunc(secureMiddleware.HandlerFuncWithNext))
	n.UseHandler(router)
	startupMessage := "===> Starting app "+ appEnv.AppName +""
	startupMessage = startupMessage + " on port " + appEnv.Port
	startupMessage = startupMessage + " in " + appEnv.Server + " mode."
	appEnv.Logger.Println(startupMessage)

	if appEnv.Server == LOCAL {
		n.Run("localhost:" + appEnv.Port)
	} else {
		n.Run(":" + appEnv.Port)
	}
}
