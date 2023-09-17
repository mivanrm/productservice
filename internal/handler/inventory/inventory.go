package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	inventoryentity "github.com/mivanrm/productservice/internal/entity/inventory"
)

type inventoryUseCase interface {
	CreateInventory(inventory *inventoryentity.Inventory) (int64, error)
	GetInventory(stockID int64) (*inventoryentity.Inventory, error)
	UpdateInventory(stockID int64, updatedInventory *inventoryentity.Inventory) error
	DeleteInventory(stockID int64) error
}
type InventoryHandler struct {
	inventoryUseCase inventoryUseCase
}

func NewInventoryHandler(inventoryUseCase inventoryUseCase) *InventoryHandler {
	return &InventoryHandler{inventoryUseCase}
}

func (h *InventoryHandler) CreateInventory(w http.ResponseWriter, r *http.Request) {
	var reqBody inventoryentity.Inventory
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Create the inventory record
	insertedID, err := h.inventoryUseCase.CreateInventory(&reqBody)
	if err != nil {
		http.Error(w, "Failed to create inventory record", http.StatusInternalServerError)
		return
	}

	// Return the inserted ID as a response
	resp := map[string]int64{"stock_id": insertedID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *InventoryHandler) GetInventory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stockID, err := strconv.ParseInt(vars["stockID"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid stock ID", http.StatusBadRequest)
		return
	}

	// Get the inventory record by stock ID
	inventoryRecord, err := h.inventoryUseCase.GetInventory(stockID)
	if err != nil {
		http.Error(w, "Inventory record not found", http.StatusNotFound)
		return
	}

	// Return the inventory record as a response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventoryRecord)
}
