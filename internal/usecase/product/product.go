package product

import (
	"errors"

	"github.com/mivanrm/productservice/internal/entity/product"
)

type productRepo interface {
	CreateProduct(product *product.Product) error
	GetProduct(productID int64) (*product.Product, error)
	UpdateProduct(productID int64, updatedProduct *product.Product) error
	DeleteProduct(productID int64) error
}

type productUsecase struct {
	productRepo productRepo
}

func New(pr productRepo) productUsecase {
	return productUsecase{
		productRepo: pr,
	}
}

func (uc *productUsecase) CreateProduct(product *product.Product) error {
	return uc.productRepo.CreateProduct(product)
}

func (uc *productUsecase) GetProduct(productID int64) (*product.Product, error) {
	if productID <= 0 {
		return nil, errors.New("invalid product ID")
	}
	return uc.productRepo.GetProduct(productID)
}

func (uc *productUsecase) UpdateProduct(productID int64, updatedProduct *product.Product) error {
	if productID <= 0 {
		return errors.New("invalid product ID")
	}
	return uc.productRepo.UpdateProduct(productID, updatedProduct)
}

func (uc *productUsecase) DeleteProduct(productID int64) error {
	if productID <= 0 {
		return errors.New("invalid product ID")
	}
	return uc.productRepo.DeleteProduct(productID)
}
