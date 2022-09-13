package routes

import (
	"github.com/arifintahu/go-rest-api/app/controllers"
	"github.com/arifintahu/go-rest-api/app/utils/handlers"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc handlers.HandlerFunc
}

type Routes []Route

var RouteList = Routes{
	Route{"GET", "/", controllers.HealthCheck},

	Route{"POST", "/books", controllers.AddBook},
	Route{"GET", "/books", controllers.ListBooks},
	Route{"GET", "/books/:id", controllers.GetBook},
	Route{"PUT", "/books/:id", controllers.UpdateBook},
	Route{"DELETE", "/books/:id", controllers.DeleteBook},
}
