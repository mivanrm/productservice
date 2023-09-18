package product

import (
	"errors"
	"fmt"

	inventoryentity "github.com/mivanrm/productservice/internal/entity/inventory"
	"github.com/mivanrm/productservice/internal/entity/product"
	productentity "github.com/mivanrm/productservice/internal/entity/product"
)

type productRepo interface {
	CreateProduct(product productentity.Product) (int64, error)
	GetProduct(productID int64) (*productentity.Product, error)
	UpdateProduct(productID int64, updatedProduct productentity.Product) error
	DeleteProduct(productID int64) error
}
type variantRepo interface {
	CreateVariant(variant *product.Variant) (int64, error)
	GetVariants(productID int64) ([]product.Variant, error)
	UpdateVariant(variantID int64, updatedVariant *product.Variant) error
	DeleteVariant(variantID int64) error
}

type inventoryRepo interface {
	CreateInventory(inventory *inventoryentity.Inventory) (int64, error)
}
type productUsecase struct {
	productRepo   productRepo
	variantRepo   variantRepo
	inventoryRepo inventoryRepo
}

func New(pr productRepo, vr variantRepo, ir inventoryRepo) productUsecase {
	return productUsecase{
		productRepo:   pr,
		variantRepo:   vr,
		inventoryRepo: ir,
	}
}

func (uc *productUsecase) CreateProduct(createproductParam productentity.CreateProductParam) error {

	productID, err := uc.productRepo.CreateProduct(productentity.Product{
		Name:        createproductParam.Name,
		Price:       createproductParam.Price,
		Description: createproductParam.Description,
	})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for _, variant := range createproductParam.Variants {
		variantID, err := uc.variantRepo.CreateVariant(&productentity.Variant{
			ParentID:   productID,
			Price:      variant.Price,
			OptionName: variant.OptionName,
			Image:      variant.Image,
		})
		if err != nil {
			fmt.Println(err.Error(), variantID)
			return err
		}
		_, err = uc.inventoryRepo.CreateInventory(&inventoryentity.Inventory{
			VariantID: variantID,
			Amount:    variant.VariantStock,
		})
		if err != nil {
			fmt.Println(err.Error(), variantID)
			return err
		}
	}

	return err
}

func (uc *productUsecase) GetProduct(productID int64) (productentity.ProductResponse, error) {

	if productID <= 0 {
		return productentity.ProductResponse{}, errors.New("invalid product ID")
	}
	imageArray := []string{}
	product, err := uc.productRepo.GetProduct(productID)
	if err != nil {
		return productentity.ProductResponse{}, nil
	}
	imageArray = append(imageArray, product.Image)

	variant, err := uc.variantRepo.GetVariants(productID)
	if err != nil {
		return productentity.ProductResponse{}, nil
	}
	return productentity.ProductResponse{
		Product:  *product,
		Image:    imageArray,
		Variants: variant,
	}, nil
}

func (uc *productUsecase) UpdateProduct(updateProductParam productentity.UpdateProductParam) error {
	err := uc.productRepo.UpdateProduct(updateProductParam.ID, productentity.Product{
		Name:        updateProductParam.Name,
		Price:       updateProductParam.Price,
		Description: updateProductParam.Description,
	})
	if err != nil {
		return err
	}
	for _, variant := range updateProductParam.Variants {
		uc.variantRepo.UpdateVariant(variant.ID, &productentity.Variant{
			Price:      variant.Price,
			OptionName: variant.OptionName,
			Image:      variant.Image,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (uc *productUsecase) DeleteProduct(productID int64) error {
	if productID <= 0 {
		return errors.New("invalid product ID")
	}
	return uc.productRepo.DeleteProduct(productID)
}
