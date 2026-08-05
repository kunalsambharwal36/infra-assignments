package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kunalsambharwal36/infra-assignments/internal/domain"
	"github.com/kunalsambharwal36/infra-assignments/internal/service"
)

type Handler struct {
	configService *service.ConfigService
}

func NewHandler(configService *service.ConfigService) *Handler {
	return &Handler{
		configService: configService,
	}
}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func (h *Handler) CreateConfig(w http.ResponseWriter, r *http.Request) {

	var config domain.Config

	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = h.configService.CreateConfig(config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Config received successfully"))
}

func (h *Handler) GetConfig(w http.ResponseWriter, r *http.Request) {

	// Read URL parameter
	vars := mux.Vars(r)

	id := vars["id"]

	config, err := h.configService.GetConfig(id)
	if err != nil {
		http.Error(w, "Config not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(config)
}
