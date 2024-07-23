package main

import (
	"log"
	"net/http"

	"github.com/JoaoGumiero/OngMais/apis"
	"github.com/JoaoGumiero/OngMais/config"
	"github.com/JoaoGumiero/OngMais/firebase"
	"github.com/JoaoGumiero/OngMais/services"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	Configs, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading enviroments: %v", err)
	}

	// Initialize Firebase
	Client := firebase.InitFirebase(*Configs)

	println(Client)
	defer func() {
		if Client != nil {
			println("diff than")
			Client.Close()
		}
	}()

	if Client == nil {
		log.Fatalf("Firestore Client is not initialized")
	}

	//Sending the cliente to handlers
	apis.RecieveFirebaseClient(Client)

	// Criar API
	states, err := services.FetchStates()
	if err != nil {
		log.Fatalf("Error fetching states: %v", err)
	}
	firebase.StoreStates(states, Client)

	cities, err := services.FetchCities()
	if err != nil {
		log.Fatalf("Error fetching cities: %v", err)
	}
	firebase.StoreCities(cities, Client)

	router := mux.NewRouter()
	apis.RegisterRoutes(router) // Register the routes defined in routes.go

	// Set up the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
