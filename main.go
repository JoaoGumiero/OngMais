package main

import (
	"log"
	"net/http"

	"github.com/JoaoGumiero/OngMais/apis"
	"github.com/JoaoGumiero/OngMais/config"
	"github.com/JoaoGumiero/OngMais/firebase"

	// "github.com/JoaoGumiero/OngMais/services" testing, tirar para n ficar dand dowload toda vez
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
	//apis.RecieveFirebaseClient(Client)

	// services.CronStateCityFetch(Configs) testing, tirar para n ficar dand dowload toda vez

	vRepo := firebase.NewVoluntaryRepository(Client)
	lRepo := firebase.NewLocationRepository(Client)

	router := mux.NewRouter()
	apis.RegisterRoutes(router, vRepo, lRepo) // Register the routes defined in routes.go

	// Set up the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
