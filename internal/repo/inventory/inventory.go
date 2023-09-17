package inventory

import (
	"database/sql"

	inventoryentity "github.com/mivanrm/productservice/internal/entity/inventory"

	"github.com/jmoiron/sqlx"
)

type inventoryRepo struct {
	db *sqlx.DB
}

func NewInventoryRepository(db *sqlx.DB) inventoryRepo {
	return inventoryRepo{db}
}

func (r *inventoryRepo) CreateInventory(inventory *inventoryentity.Inventory) (int64, error) {
	query := "INSERT INTO inventory (variant_id, amount) VALUES (?, ?)"
	result, err := r.db.Exec(query, inventory.VariantID, inventory.Amount)
	if err != nil {
		return 0, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func (r *inventoryRepo) GetInventory(stockID int64) (*inventoryentity.Inventory, error) {
	query := "SELECT * FROM inventory WHERE stock_id = ?"
	var inventory inventoryentity.Inventory
	err := r.db.Get(&inventory, query, stockID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return &inventory, nil
}

func (r *inventoryRepo) UpdateInventory(variantID int64, updatedInventory *inventoryentity.Inventory) error {
	query := "UPDATE inventory SET amount = ? WHERE variant_id = ?"
	_, err := r.db.Exec(query, updatedInventory.Amount, variantID)
	if err != nil {
		return err
	}
	return nil
}
