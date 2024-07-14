package main

import (
	"log"
	"net/http"

	"github.com/JoaoGumiero/OngMais/apis"
	"github.com/JoaoGumiero/OngMais/firebase"
	"github.com/JoaoGumiero/OngMais/services"
	"github.com/gorilla/mux"
)

func main() {
	client := firebase.InitFirebase()
	println(client)
	defer func() {
		if client != nil {
			println("diff than")
			client.Close()
		}
	}()

	if client == nil {
		log.Fatalf("Firestore client is not initialized")
	}

	states, err := services.FetchStates()
	if err != nil {
		log.Fatalf("Error fetching states: %v", err)
	}
	firebase.StoreStates(states, client)

	cities, err := services.FetchCities()
	if err != nil {
		log.Fatalf("Error fetching cities: %v", err)
	}
	firebase.StoreCities(cities, client)

	router := mux.NewRouter()
	apis.RegisterRoutes(router) // Register the routes defined in routes.go

	// Set up the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
