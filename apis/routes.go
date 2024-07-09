package apis

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	// Register routes for the voluntary entity
	router.HandleFunc("/voluntaries", CreateVoluntary).Methods("POST")
	router.HandleFunc("/voluntaries/{id}", GetVoluntary).Methods("GET")
	router.HandleFunc("/voluntaries/{id}", UpdateVoluntary).Methods("PUT")
	router.HandleFunc("/voluntaries/{id}", DeleteVoluntary).Methods("DELETE")
	router.HandleFunc("/states", getStates).Methods("GET")
	router.HandleFunc("/cities", getCities).Methods("GET")
}
