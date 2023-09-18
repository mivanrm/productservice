package product

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/mivanrm/productservice/internal/entity/product"
	"github.com/stretchr/testify/assert"
)

func TestProductRepo_CreateProduct(t *testing.T) {
	// Create a new mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	// Create an sqlx.DB instance using the mock database connection
	db := sqlx.NewDb(mockDB, "sqlmock")

	// Create a product repository
	repo := New(db)

	// Mock the expected database query and result
	mock.ExpectQuery(`INSERT INTO products (.+) RETURNING product_id`).
		WithArgs("Test Product", "Test Description", "test.jpg", 10.99).
		WillReturnRows(sqlmock.NewRows([]string{"product_id"}).AddRow(1))

	// Create a sample product to insert
	productToInsert := product.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Image:       "test.jpg",
		Price:       10.99,
	}

	// Call the CreateProduct function
	insertedID, err := repo.CreateProduct(productToInsert)
	if err != nil {
		t.Fatalf("Error creating product: %v", err)
	}

	// Verify that the expected ID was returned
	assert.Equal(t, int64(1), insertedID)
}

func TestProductRepo_GetProduct(t *testing.T) {
	// Create a new mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	// Create an sqlx.DB instance using the mock database connection
	db := sqlx.NewDb(mockDB, "sqlmock")

	// Create a product repository
	repo := New(db)

	// Mock the expected database query and result
	expectedProduct := product.Product{
		ID:          1,
		Name:        "Test Product",
		Description: "Test Description",
		Image:       "test.jpg",
		Price:       10.99,
	}

	rows := sqlmock.NewRows([]string{"product_id", "name", "description", "image", "price"}).
		AddRow(expectedProduct.ID, expectedProduct.Name, expectedProduct.Description, expectedProduct.Image, expectedProduct.Price)

	mock.ExpectQuery(`SELECT \* FROM products WHERE product_id = \$1`).
		WithArgs(1).
		WillReturnRows(rows)

	// Call the GetProduct function
	productByID, err := repo.GetProduct(1)
	if err != nil {
		t.Fatalf("Error getting product by ID: %v", err)
	}

	// Verify that the expected product was returned
	assert.Equal(t, &expectedProduct, productByID)
}

func TestProductRepo_UpdateProduct(t *testing.T) {
	// Create a new mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	// Create an sqlx.DB instance using the mock database connection
	db := sqlx.NewDb(mockDB, "sqlmock")

	// Create a product repository
	repo := New(db)

	// Mock the expected database query and result
	mock.ExpectExec(`UPDATE products SET name=\$1, description=\$2, price=\$3 WHERE product_id=\$4`).
		WithArgs("Updated Product", "Updated Description", 19.99, 1).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row affected

	// Create an updated product
	updatedProduct := product.Product{
		Name:        "Updated Product",
		Description: "Updated Description",
		Price:       19.99,
	}

	// Call the UpdateProduct function
	err = repo.UpdateProduct(1, updatedProduct)
	if err != nil {
		t.Fatalf("Error updating product: %v", err)
	}

	// Verify that the product was updated
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepo_DeleteProduct(t *testing.T) {
	// Create a new mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	// Create an sqlx.DB instance using the mock database connection
	db := sqlx.NewDb(mockDB, "sqlmock")

	// Create a product repository
	repo := New(db)

	// Mock the expected database query and result for deleting a product
	mock.ExpectExec(`DELETE FROM products WHERE product_id = \$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row affected

	// Call the DeleteProduct function
	err = repo.DeleteProduct(1)
	if err != nil {
		t.Fatalf("Error deleting product: %v", err)
	}

	// Verify that the product was deleted
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepo_UpdateProductRating(t *testing.T) {
	// Create a new mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	// Create an sqlx.DB instance using the mock database connection
	db := sqlx.NewDb(mockDB, "sqlmock")

	// Create a product repository
	repo := New(db)

	// Mock the expected database query and result for updating product rating
	mock.ExpectExec(`UPDATE products SET rating=\$1, rating_count=\$2 WHERE product_id=\$3`).
		WithArgs(4.5, 100, 1).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row affected

	// Call the UpdateProductRating function
	err = repo.UpdateProductRating(1, 4.5, 100)
	if err != nil {
		t.Fatalf("Error updating product rating: %v", err)
	}

	// Verify that the product rating was updated
	assert.NoError(t, mock.ExpectationsWereMet())
}
