package main

import (
	"log"
	"net/http"

	"github.com/JoaoGumiero/OngMais/apis"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	apis.RegisterRoutes(router) // Register the routes defined in routes.go

	// Set up the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
