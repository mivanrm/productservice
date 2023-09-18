package usecase

import (
	"github.com/mivanrm/productservice/internal/entity/review"
	reviewentity "github.com/mivanrm/productservice/internal/entity/review"
)

type reviewRepository interface {
	CreateReview(review *reviewentity.Review) (int64, error)
	GetReview(reviewID int64) (*reviewentity.Review, error)
	UpdateReview(reviewID int64, updatedReview *reviewentity.Review) error
	DeleteReview(reviewID int64) error
}
type reviewUseCase struct {
	reviewRepo reviewRepository
}

func NewReviewUseCase(reviewRepo reviewRepository) reviewUseCase {
	return reviewUseCase{reviewRepo: reviewRepo}
}

func (uc *reviewUseCase) CreateReview(review *review.Review) (int64, error) {
	return uc.reviewRepo.CreateReview(review)
}

func (uc *reviewUseCase) GetReview(reviewID int64) (*review.Review, error) {
	return uc.reviewRepo.GetReview(reviewID)
}

func (uc *reviewUseCase) UpdateReview(reviewID int64, updatedReview *review.Review) error {
	err := uc.reviewRepo.UpdateReview(reviewID, updatedReview)
	if err != nil {
		return err
	}
	return nil
}

func (uc *reviewUseCase) DeleteReview(reviewID int64) error {
	err := uc.reviewRepo.DeleteReview(reviewID)
	if err != nil {
		return err
	}
	return nil
}
