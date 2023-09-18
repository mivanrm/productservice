package product

import (
	"github.com/jmoiron/sqlx"
	"github.com/mivanrm/productservice/internal/entity/product"
)

type productRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) productRepo {
	return productRepo{
		db: db,
	}
}

// Create a new product
func (pr *productRepo) CreateProduct(product product.Product) (int64, error) {
	query := "INSERT INTO products (name, description, image, price) VALUES ($1, $2, $3, $4) RETURNING product_id "
	var insertedID int64
	err := pr.db.QueryRow(query, product.Name, product.Description, product.Image, product.Price).Scan(&insertedID)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
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
func (pr *productRepo) UpdateProduct(productID int64, updatedProduct product.Product) error {
	query := "UPDATE products SET name=$1, description=$2, price=$3 WHERE product_id=$4"
	_, err := pr.db.Exec(query, updatedProduct.Name, updatedProduct.Description, updatedProduct.Price, productID)
	return err
}

// Delete a product
func (pr *productRepo) DeleteProduct(productID int64) error {
	query := "DELETE FROM products WHERE product_id = $1"
	_, err := pr.db.Exec(query, productID)
	return err
}

// Update a product rating
func (pr *productRepo) UpdateProductRating(productID int64, rating float64, rating_count int64) error {
	query := "UPDATE products SET rating=$1, rating_count=$2 WHERE product_id=$3"
	_, err := pr.db.Exec(query, rating, rating_count, productID)

	return err
}
