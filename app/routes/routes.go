package routes

import "github.com/arifintahu/go-rest-api/app/handlers"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc handlers.HandlerFunc
}

type Routes []Route

var RouteList = Routes{
	Route{"HealthCheck", "GET", "/", handlers.HealthCheck},

	Route{"AddBook", "POST", "/books", handlers.AddBook},
	Route{"ListBooks", "GET", "/books", handlers.ListBooks},
	Route{"GetBook", "GET", "/books/{id}", handlers.GetBook},
	Route{"UpdateBook", "PUT", "/books/{id}", handlers.UpdateBook},
	Route{"DeleteBook", "DELETE", "/books/{id}", handlers.DeleteBook},
}
