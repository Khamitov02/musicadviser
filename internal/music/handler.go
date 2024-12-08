package music

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
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
		r.Post("/api/v1/putMusic", h.putMusic)
	})
}

func (h *Handler) getProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Touched handler GetProducts")
	// validate r
	data, err := h.service.Products(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "", data)
}

func (h *Handler) putMusic(w http.ResponseWriter, r *http.Request) {
	log.Println("Touched handler putMusic")

}
