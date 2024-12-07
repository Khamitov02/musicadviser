package fridge

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	router  *chi.Mux
	service Service
}

func NewHandler(router *chi.Mux, service Service) *Handler {
	return &Handler{
		router:  router,
		service: service,
	}
}

func (h *Handler) Register() {
	h.router.Group(func(r chi.Router) {
		r.Get("/api/v1/products", h.getProducts)
		r.Post("/api/v1/products", h.postProducts)
	})
}

func (h *Handler) getProducts(w http.ResponseWriter, r *http.Request) {
	// validate r
	data, err := h.service.Products(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "", data)
}

func (h *Handler) postProducts(w http.ResponseWriter, r *http.Request) {

}
