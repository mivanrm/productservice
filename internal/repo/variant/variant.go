package variant

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/mivanrm/productservice/internal/entity/product"
)

type variantRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) variantRepo {
	return variantRepo{db}
}

func (r *variantRepo) CreateVariant(variant *product.Variant) (int64, error) {
	query := "INSERT INTO variants (option_name, price, parent_id) VALUES ($1, $2, $3) RETURNING variant_id"
	var insertedID int64
	err := r.db.QueryRow(query, variant.OptionName, variant.Price, variant.ParentID).Scan(&insertedID)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}
func (r *variantRepo) GetVariants(productID int64) ([]product.Variant, error) {
	query := "SELECT * FROM variants WHERE parent_id = $1"
	response := []product.Variant{}
	err := r.db.Select(&response, query, productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return response, nil
}

func (r *variantRepo) UpdateVariant(variantID int64, updatedVariant *product.Variant) error {
	query := "UPDATE variants SET option_name = $1, price = $2 WHERE variant_id = $3"
	_, err := r.db.Exec(query, updatedVariant.OptionName, updatedVariant.Price, variantID)
	if err != nil {
		return err
	}
	return nil
}

func (r *variantRepo) DeleteVariant(variantID int64) error {
	query := "DELETE FROM variants WHERE variant_id = $1"
	_, err := r.db.Exec(query, variantID)
	if err != nil {
		return err
	}
	return nil
}
