package apis

// Make the handler call the service layer files and the make the business filter logic there.
import (
	"encoding/json"
	"net/http"

	"github.com/JoaoGumiero/OngMais/entities"
	"github.com/JoaoGumiero/OngMais/services"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type VoluntaryHandler struct {
	voluntaryService *services.VoluntaryService
}

func NewVoluntaryHandler(voluntaryService *services.VoluntaryService) *VoluntaryHandler {
	return &VoluntaryHandler{voluntaryService: voluntaryService}
}

var validate = validator.New()

func (h *VoluntaryHandler) CreateVoluntary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var voluntary entities.Voluntary
	if err := json.NewDecoder(r.Body).Decode(&voluntary); err != nil {
		http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.voluntaryService.AddVoluntaryService(&voluntary, ctx); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(voluntary)
	w.WriteHeader(http.StatusCreated)
}

// GetVoluntary retrieves a voluntary by ID
func (h *VoluntaryHandler) GetVoluntaryById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"] // Using mux to extract variables from route
	voluntary, err := h.voluntaryService.GetVoluntaryByIdService(id, ctx)
	if err != nil {
		http.Error(w, "Failed to fetch voluntary: "+err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(voluntary)
}

// UpdateVoluntary updates an existing voluntary
func (h *VoluntaryHandler) UpdateVoluntary(w http.ResponseWriter, r *http.Request) {
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
	err := h.voluntaryService.UpdateVoluntaryService(id, updates, ctx)
	if err != nil {
		http.Error(w, "Failed to update voluntary: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteVoluntary removes a voluntary from the database
func (h *VoluntaryHandler) DeleteVoluntary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]
	err := h.voluntaryService.DeleteVoluntaryService(id, ctx)
	if err != nil {
		http.Error(w, "Failed to fetch voluntary: "+err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
