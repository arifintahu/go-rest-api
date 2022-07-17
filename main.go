package main

import (
	"go-rest-api-2/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func main() {
	logger := log.New(os.Stdout, "test-api", log.LstdFlags)
	render := render.New()
	sm := mux.NewRouter().StrictSlash(true)
	handler := handlers.InitHandler(logger, render)

	sm.HandleFunc("/", handler.PingConnection)
	sm.HandleFunc("/books", handler.ListBooks).Methods("GET")
	sm.HandleFunc("/books/{id}", handler.GetBook).Methods("GET")

	server := &http.Server{
		Addr:         ":3000",
		Handler:      sm,
	}

	log.Fatal(server.ListenAndServe())
}
