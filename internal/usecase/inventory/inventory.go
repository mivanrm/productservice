package usecase

import (
	inventoryentity "github.com/mivanrm/productservice/internal/entity/inventory"
)

type inventoryRepository interface {
	CreateInventory(inventory inventoryentity.Inventory) (int64, error)
	GetInventory(stockID int64) (inventoryentity.Inventory, error)
	UpdateInventory(stockID int64, updatedInventory *inventoryentity.Inventory) error
	DeleteInventory(stockID int64) error
}

type inventoryUseCase struct {
	inventoryRepo inventoryRepository
}

func NewInventoryUseCase(inventoryRepo inventoryRepository) inventoryUseCase {
	return inventoryUseCase{inventoryRepo}
}

func (uc *inventoryUseCase) GetInventory(VariantID int64) (inventoryentity.Inventory, error) {
	return uc.inventoryRepo.GetInventory(VariantID)
}

func (uc *inventoryUseCase) UpdateInventory(VariantID int64, updatedInventory *inventoryentity.Inventory) error {
	return uc.inventoryRepo.UpdateInventory(VariantID, updatedInventory)
}
