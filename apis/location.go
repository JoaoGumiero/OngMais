package apis

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/JoaoGumiero/OngMais/services"
)

type LocationHandler struct {
	LocationService *services.LocationService
}

func NewLocationHandler(LocationService *services.LocationService) *LocationHandler {
	return &LocationHandler{LocationService: LocationService}
}

// Get all Brazilian States
func (h *LocationHandler) getStates(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	states, err := h.LocationService.GetStatesService(ctx)
	if err != nil {
		http.Error(w, "Falied to get all States: "+err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(states)
}

// Get all Brazilian Cities
func (h *LocationHandler) getCities(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cities, err := h.LocationService.GetCitiesService(ctx)
	if err != nil {
		http.Error(w, "Falied to get all Cities: "+err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(cities)
}
