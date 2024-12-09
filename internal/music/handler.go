package music

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	router  *chi.Mux
	service Service
}

type MusicRequest struct {
	UserID string   `json:"user_id"`
	Bands  []string `json:"bands"`
}

func NewHandler(router *chi.Mux, service Service) *Handler {
	return &Handler{
		router:  router,
		service: service,
	}
}

func (h *Handler) Register() {
	h.router.Group(func(r chi.Router) {
		r.Get("/api/v1/getMusic", h.getMusic)
		r.Post("/api/v1/putMusic", h.putMusic)
	})
}

func (h *Handler) getMusic(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler: GetMusic accessed - UserAgent: %s, RemoteAddr: %s", r.UserAgent(), r.RemoteAddr)
	
	userBands, err := h.service.GetAllUserBands(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(userBands); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) putMusic(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler: PutMusic accessed - UserAgent: %s, RemoteAddr: %s", r.UserAgent(), r.RemoteAddr)
	var req MusicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Process each band in the request
	for _, bandName := range req.Bands {
		product := Product{
			UserID:   req.UserID,
			BandName: bandName,
		}
		
		_, err := h.service.Place(r.Context(), product)
		if err != nil {
			// If the band already exists, continue to the next one
			if strings.Contains(err.Error(), "already exists") {
				continue
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Music bands added successfully")
}
