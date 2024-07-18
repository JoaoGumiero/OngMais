package apis

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	// Register routes for the voluntary entity
	router.HandleFunc("/voluntary", CreateVoluntary).Methods("POST")
	router.HandleFunc("/voluntary/{id}", GetVoluntary).Methods("GET")
	router.HandleFunc("/voluntary/{id}", UpdateVoluntary).Methods("PUT")
	router.HandleFunc("/voluntary/{id}", DeleteVoluntary).Methods("DELETE")
	router.HandleFunc("/states", getStates).Methods("GET")
	router.HandleFunc("/cities", getCities).Methods("GET")
}
