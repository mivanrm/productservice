package product

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mivanrm/productservice/internal/entity/product"
)

// type Repo interface {
// 	CreateProduct(product *product.Product) error
// 	GetProduct(productID int64) (*product.Product, error)
// 	UpdateProduct(productID int64, updatedProduct *product.Product) error
// 	DeleteProduct(productID int64) error
// }

type productRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) productRepo {
	return productRepo{
		db: db,
	}
}

// Create a new product
func (pr *productRepo) CreateProduct(product *product.Product) error {
	query := "INSERT INTO product (name, description, rating) VALUES (?, ?, ?)"

	result, err := pr.db.Exec(query, product.Name, product.Description, product.Rating)
	fmt.Println(result)
	if err != nil {
		return err
	}
	return nil
}

// Get a product by ID
func (pr *productRepo) GetProduct(productID int64) (*product.Product, error) {
	query := "SELECT * FROM products WHERE product_id = $1"
	var product product.Product
	err := pr.db.Get(&product, query, productID)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Update a product
func (pr *productRepo) UpdateProduct(productID int64, updatedProduct *product.Product) error {
	query := "UPDATE products SET name=$1, description=$2, rating=$3 WHERE product_id=$4"
	_, err := pr.db.Exec(query, updatedProduct.Name, updatedProduct.Description, updatedProduct.Rating, productID)
	return err
}

// Delete a product
func (pr *productRepo) DeleteProduct(productID int64) error {
	query := "DELETE FROM products WHERE product_id = $1"
	_, err := pr.db.Exec(query, productID)
	return err
}
