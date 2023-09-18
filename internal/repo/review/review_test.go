package review

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	reviewmodel "github.com/mivanrm/productservice/internal/entity/review"
	"github.com/stretchr/testify/assert"
)

func TestReviewRepo_CreateReview(t *testing.T) {
	// Create a new mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	// Create an sqlx.DB instance using the mock database connection
	db := sqlx.NewDb(mockDB, "sqlmock")

	// Create a review repository
	repo := New(db)

	// Mock the expected database query and result
	mock.ExpectQuery(`INSERT INTO reviews (.+) RETURNING review_id`).
		WithArgs(1, "Test Review", 4).
		WillReturnRows(sqlmock.NewRows([]string{"review_id"}).AddRow(1))

	// Create a sample review to insert
	reviewToInsert := &reviewmodel.Review{
		ProductID:  1,
		ReviewText: "Test Review",
		Rating:     4,
	}

	// Call the CreateReview function
	insertedID, err := repo.CreateReview(reviewToInsert)
	if err != nil {
		t.Fatalf("Error creating review: %v", err)
	}

	// Verify that the expected ID was returned
	assert.Equal(t, int64(1), insertedID)
}

func TestReviewRepo_GetReview(t *testing.T) {
	// Create a new mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	// Create an sqlx.DB instance using the mock database connection
	db := sqlx.NewDb(mockDB, "sqlmock")

	// Create a review repository
	repo := New(db)

	// Mock the expected database query and result
	expectedReview := &reviewmodel.Review{
		ReviewID:   1,
		ProductID:  1,
		ReviewText: "Test Review",
		Rating:     4,
	}

	rows := sqlmock.NewRows([]string{"review_id", "product_id", "review_text", "rating"}).
		AddRow(expectedReview.ReviewID, expectedReview.ProductID, expectedReview.ReviewText, expectedReview.Rating)

	mock.ExpectQuery(`SELECT \* FROM reviews WHERE review_id = \$1`).
		WithArgs(1).
		WillReturnRows(rows)

	// Call the GetReview function
	reviewByID, err := repo.GetReview(1)
	if err != nil {
		t.Fatalf("Error getting review by ID: %v", err)
	}

	// Verify that the expected review was returned
	assert.Equal(t, expectedReview, reviewByID)
}

func TestReviewRepo_UpdateReview(t *testing.T) {
	// Create a new mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	// Create an sqlx.DB instance using the mock database connection
	db := sqlx.NewDb(mockDB, "sqlmock")

	// Create a review repository
	repo := New(db)

	// Mock the expected database query and result for updating a review
	mock.ExpectExec(`UPDATE reviews SET review_text = \$1, rating = \$2 WHERE review_id = \$3`).
		WithArgs("Updated Review", 3, 1).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row affected

	// Create an updated review
	updatedReview := &reviewmodel.Review{
		ReviewText: "Updated Review",
		Rating:     3,
	}

	// Call the UpdateReview function
	err = repo.UpdateReview(1, updatedReview)
	if err != nil {
		t.Fatalf("Error updating review: %v", err)
	}

	// Verify that the review was updated
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestReviewRepo_DeleteReview(t *testing.T) {
	// Create a new mock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer mockDB.Close()

	// Create an sqlx.DB instance using the mock database connection
	db := sqlx.NewDb(mockDB, "sqlmock")

	// Create a review repository
	repo := New(db)

	// Mock the expected database query and result for deleting a review
	mock.ExpectExec(`DELETE FROM reviews WHERE review_id = \$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row affected

	// Call the DeleteReview function
	err = repo.DeleteReview(1)
	if err != nil {
		t.Fatalf("Error deleting review: %v", err)
	}

	// Verify that the review was deleted
	assert.NoError(t, mock.ExpectationsWereMet())
}
