package review

import (
	"fmt"
	"net/http"
)

type Handler interface {
	AddReview(w http.ResponseWriter, r *http.Request)
	UpdateReview(w http.ResponseWriter, r *http.Request)
	DeleteReview(w http.ResponseWriter, r *http.Request)
	GetReview(w http.ResponseWriter, r *http.Request)
}

type handler struct {
}

func New() Handler {
	return &handler{}
}

func (h *handler) AddReview(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", "category")
}
func (h *handler) UpdateReview(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", "category")
}
func (h *handler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", "category")
}
func (h *handler) GetReview(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", "category")
}
