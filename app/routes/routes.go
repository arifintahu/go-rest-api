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
	Route{"PingConnection", "GET", "/", handlers.HealthCheck},
	Route{"ListBooks", "GET", "/books", handlers.ListBooks},
	Route{"GetBook", "GET", "/books/{id}", handlers.GetBook},
}
