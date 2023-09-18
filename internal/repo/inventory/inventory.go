package inventory

import (
	"database/sql"

	"github.com/lib/pq"
	inventoryentity "github.com/mivanrm/productservice/internal/entity/inventory"

	"github.com/jmoiron/sqlx"
)

type inventoryRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) inventoryRepo {
	return inventoryRepo{db}
}
func (r *inventoryRepo) CreateInventory(inventory inventoryentity.Inventory) (int64, error) {
	query := "INSERT INTO inventory (variant_id, amount) VALUES ($1, $2) RETURNING stock_id"
	var insertedID int64
	err := r.db.QueryRow(query, inventory.VariantID, inventory.Amount).Scan(&insertedID)
	if err != nil {
		return 0, err
	}
	return insertedID, nil
}

func (r *inventoryRepo) GetInventory(stockID int64) (inventoryentity.Inventory, error) {
	query := "SELECT * FROM inventory WHERE stock_id = $1"
	var inventory inventoryentity.Inventory
	err := r.db.Get(&inventory, query, stockID)
	if err != nil {
		if err == sql.ErrNoRows {
			return inventoryentity.Inventory{}, sql.ErrNoRows
		}
		return inventoryentity.Inventory{}, err
	}
	return inventory, nil
}

func (r *inventoryRepo) UpdateInventory(variantID int64, updatedInventory *inventoryentity.Inventory) error {
	query := "UPDATE inventory SET amount = $1 WHERE variant_id = $2"
	_, err := r.db.Exec(query, updatedInventory.Amount, variantID)
	if err != nil {
		return err
	}
	return nil
}

func (r *inventoryRepo) GetInventoryByVariantIDs(variantIDs []int64) (map[int64]int64, error) {
	query := "SELECT variant_id, amount FROM inventory WHERE variant_id = ANY($1)"
	stockMap := make(map[int64]int64)
	rows, err := r.db.Queryx(query, pq.Array(variantIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var variantID int64
		var stock int64
		err := rows.Scan(&variantID, &stock)
		if err != nil {
			return nil, err
		}
		stockMap[variantID] = stock
	}

	return stockMap, nil
}
