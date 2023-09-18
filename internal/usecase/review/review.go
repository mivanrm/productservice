package usecase

import (
	productentity "github.com/mivanrm/productservice/internal/entity/product"
	reviewentity "github.com/mivanrm/productservice/internal/entity/review"
)

type productRepo interface {
	GetProduct(productID int64) (*productentity.Product, error)
	UpdateProductRating(productID int64, rating float64, rating_count int64) error
}

type reviewRepository interface {
	CreateReview(review *reviewentity.Review) (int64, error)
	GetReview(reviewID int64) (*reviewentity.Review, error)
	UpdateReview(reviewID int64, updatedReview *reviewentity.Review) error
	DeleteReview(reviewID int64) error
}
type reviewUseCase struct {
	reviewRepo  reviewRepository
	productRepo productRepo
}

func New(reviewRepo reviewRepository, productRepo productRepo) reviewUseCase {
	return reviewUseCase{reviewRepo: reviewRepo, productRepo: productRepo}
}

func (uc *reviewUseCase) CreateReview(review *reviewentity.Review) (int64, error) {
	var rating float64
	product, err := uc.productRepo.GetProduct(review.ProductID)
	if err != nil {
		return 0, err
	}
	if product.Rating == 0 {
		rating += float64(review.Rating)
	} else {
		rating = float64(((float32(product.Rating) * float32(product.RatingCount)) + float32(review.Rating)) / float32(product.RatingCount+1))
	}
	err = uc.productRepo.UpdateProductRating(review.ProductID, float64(rating), product.RatingCount+1)
	if err != nil {
		return 0, err
	}
	insertID, err := uc.reviewRepo.CreateReview(review)
	if err != nil {
		return 0, err
	}
	return insertID, nil
}

func (uc *reviewUseCase) GetReview(reviewID int64) (*reviewentity.Review, error) {
	return uc.reviewRepo.GetReview(reviewID)
}

func (uc *reviewUseCase) UpdateReview(reviewID int64, updatedReview *reviewentity.Review) error {
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
