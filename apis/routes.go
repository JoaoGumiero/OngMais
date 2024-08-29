package apis

import (
	"github.com/JoaoGumiero/OngMais/firebase"
	"github.com/JoaoGumiero/OngMais/services"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, vRepo *firebase.VoluntaryRepository, lRepo *firebase.LocationRepository) {
	// Initialize Services
	voluntaryService := services.NewVoluntaryService(vRepo)
	locationService := services.NewLocationService(lRepo)
	// Initialize handlers
	voluntaryHandler := NewVoluntaryHandler(voluntaryService)
	locationHandler := NewLocationHandler(locationService)
	// Register routes for the voluntary entity
	router.HandleFunc("/voluntary", voluntaryHandler.CreateVoluntary).Methods("POST")
	router.HandleFunc("/voluntary/{id}", voluntaryHandler.GetVoluntaryById).Methods("GET")
	router.HandleFunc("/voluntary/{id}", voluntaryHandler.UpdateVoluntary).Methods("PUT")
	router.HandleFunc("/voluntary/{id}", voluntaryHandler.DeleteVoluntary).Methods("DELETE")
	router.HandleFunc("/states", locationHandler.getStates).Methods("GET")
	router.HandleFunc("/cities", locationHandler.getCities).Methods("GET")
}
