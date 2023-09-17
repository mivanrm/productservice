package review

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mivanrm/productservice/internal/entity/review"
	reviewmodel "github.com/mivanrm/productservice/internal/entity/review"
)

type reviewRepo struct {
	db *sqlx.DB
}

func NewReviewRepository(db *sqlx.DB) reviewRepo {
	return reviewRepo{db}
}

func (r *reviewRepo) CreateReview(review *reviewmodel.Review) (int64, error) {
	query := "INSERT INTO reviews (product_id, review_text, rating) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, review.ProductID, review.ReviewText, review.Rating)
	if err != nil {
		return 0, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func (r *reviewRepo) GetReview(reviewID int64) (*review.Review, error) {
	query := "SELECT * FROM reviews WHERE review_id = ?"
	var review review.Review
	err := r.db.Get(&review, query, reviewID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Review not found")
		}
		return nil, err
	}
	return &review, nil
}

func (r *reviewRepo) UpdateReview(reviewID int64, updatedReview *review.Review) error {
	query := "UPDATE reviews SET review_text = ?, rating = ? WHERE review_id = ?"
	_, err := r.db.Exec(query, updatedReview.ReviewText, updatedReview.Rating, reviewID)
	if err != nil {
		return err
	}
	return nil
}

func (r *reviewRepo) DeleteReview(reviewID int64) error {
	query := "DELETE FROM reviews WHERE review_id = ?"
	_, err := r.db.Exec(query, reviewID)
	if err != nil {
		return err
	}
	return nil
}

func (r *reviewRepo) GetReviewsByProductID(productID int64) ([]*reviewmodel.Review, error) {
	query := "SELECT * FROM reviews WHERE product_id = ?"
	var reviews []*reviewmodel.Review
	err := r.db.Select(&reviews, query, productID)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}
