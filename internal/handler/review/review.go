package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mivanrm/productservice/internal/entity/review"
)

type reviewUsecase interface {
	CreateReview(review *review.Review) (int64, error)
	GetReview(reviewID int64) (*review.Review, error)
	UpdateReview(reviewID int64, updatedReview *review.Review) error
	DeleteReview(reviewID int64) error
}
type ReviewHandler struct {
	reviewUseCase reviewUsecase
}

func New(reviewUseCase reviewUsecase) *ReviewHandler {
	return &ReviewHandler{reviewUseCase: reviewUseCase}
}

func (h *ReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	var newReview review.Review
	err := json.NewDecoder(r.Body).Decode(&newReview)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	insertedID, err := h.reviewUseCase.CreateReview(&newReview)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create review", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]int64{"review_id": insertedID}
	json.NewEncoder(w).Encode(response)
}

func (h *ReviewHandler) GetReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid review ID", http.StatusBadRequest)
		return
	}

	reviewData, err := h.reviewUseCase.GetReview(reviewID)
	if err != nil {
		http.Error(w, "Review not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviewData)
}

func (h *ReviewHandler) UpdateReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid review ID", http.StatusBadRequest)
		return
	}

	var updatedReview review.Review
	err = json.NewDecoder(r.Body).Decode(&updatedReview)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.reviewUseCase.UpdateReview(reviewID, &updatedReview)
	if err != nil {
		http.Error(w, "Failed to update review", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ReviewHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid review ID", http.StatusBadRequest)
		return
	}

	err = h.reviewUseCase.DeleteReview(reviewID)
	if err != nil {
		http.Error(w, "Failed to delete review", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
