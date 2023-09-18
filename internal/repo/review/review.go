package review

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	reviewmodel "github.com/mivanrm/productservice/internal/entity/review"
)

type reviewRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) reviewRepo {
	return reviewRepo{db}
}
func (r *reviewRepo) CreateReview(review *reviewmodel.Review) (int64, error) {
	query := "INSERT INTO reviews (product_id, review_text, rating) VALUES ($1, $2, $3) RETURNING review_id"
	var insertedID int64
	err := r.db.QueryRow(query, review.ProductID, review.ReviewText, review.Rating).Scan(&insertedID)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func (r *reviewRepo) GetReview(reviewID int64) (*reviewmodel.Review, error) {
	query := "SELECT * FROM reviews WHERE review_id = $1"
	var review reviewmodel.Review
	err := r.db.Get(&review, query, reviewID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Review not found")
		}
		return nil, err
	}
	return &review, nil
}

func (r *reviewRepo) UpdateReview(reviewID int64, updatedReview *reviewmodel.Review) error {
	query := "UPDATE reviews SET review_text = $1, rating = $2 WHERE review_id = $3"
	_, err := r.db.Exec(query, updatedReview.ReviewText, updatedReview.Rating, reviewID)
	if err != nil {
		return err
	}
	return nil
}

func (r *reviewRepo) DeleteReview(reviewID int64) error {
	query := "DELETE FROM reviews WHERE review_id = $1"
	_, err := r.db.Exec(query, reviewID)
	if err != nil {
		return err
	}
	return nil
}

func (r *reviewRepo) GetReviewsByProductID(productID int64) ([]*reviewmodel.Review, error) {
	query := "SELECT * FROM reviews WHERE product_id = $1"
	var reviews []*reviewmodel.Review
	err := r.db.Select(&reviews, query, productID)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}
