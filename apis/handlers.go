package apis

import (
	"context"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/JoaoGumiero/OngMais/entities"
	"github.com/JoaoGumiero/OngMais/firebase"
	"github.com/JoaoGumiero/OngMais/utils" // Adjust the import path as necessary
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var client *firestore.Client

func init() {
	client = firebase.InitFirebase()
}

var validate = validator.New()

func CreateVoluntary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var voluntary entities.Voluntary
	if err := json.NewDecoder(r.Body).Decode(&voluntary); err != nil {
		http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Validate email and phone number
	if err := validate.Struct(voluntary); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	docRef, _, err := client.Collection("voluntaries").Add(ctx, voluntary)
	if err != nil {
		http.Error(w, "Failed to add voluntary: "+err.Error(), http.StatusInternalServerError)
		return
	}
	voluntary.ID = docRef.ID // Optionally update the ID with the Firestore generated ID
	json.NewEncoder(w).Encode(voluntary)
	w.WriteHeader(http.StatusCreated)
}

// GetVoluntary retrieves a voluntary by ID
func GetVoluntary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"] // Using mux to extract variables from route
	doc, err := client.Collection("voluntaries").Doc(id).Get(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch voluntary: "+err.Error(), http.StatusNotFound)
		return
	}
	var voluntary entities.Voluntary
	doc.DataTo(&voluntary)
	json.NewEncoder(w).Encode(voluntary)
}

// UpdateVoluntary updates an existing voluntary
func UpdateVoluntary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Validate email and phone number
	// if email, ok := updates["email"]; ok {
	// 	emailValidator := utils.EmailValidator{Email: email.(string)}
	// 	if !emailValidator.IsValid() {
	// 		http.Error(w, "Invalid email format", http.StatusBadRequest)
	// 		return
	// 	}
	// }
	// if phone, ok := updates["phone"]; ok {
	// 	phoneValidator := utils.PhoneValidator{Phone: phone.(string)}
	// 	if !phoneValidator.IsValid() {
	// 		http.Error(w, "Invalid phone format", http.StatusBadRequest)
	// 		return
	// 	}
	// }

	// Validate updates
	if err := utils.ValidateUpdatesStruct(updates); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err := client.Collection("voluntaries").Doc(id).Set(ctx, updates, firestore.MergeAll)
	if err != nil {
		http.Error(w, "Failed to update voluntary: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteVoluntary removes a voluntary from the database
func DeleteVoluntary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]
	_, err := client.Collection("voluntaries").Doc(id).Delete(ctx)
	if err != nil {
		http.Error(w, "Failed to delete voluntary: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Get all Brazilian States
func getStates(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	iter := client.Collection("states").Documents(ctx)
	var states []entities.State
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var state entities.State
		doc.DataTo(&state)
		states = append(states, state)
	}
	json.NewEncoder(w).Encode(states)
}

// Get all Brazilian Cities
func getCities(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	iter := client.Collection("cities").Documents(ctx)
	var cities []entities.City
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var city entities.City
		doc.DataTo(&city)
		cities = append(cities, city)
	}
	json.NewEncoder(w).Encode(cities)
}
