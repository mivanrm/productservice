package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	inventoryentity "github.com/mivanrm/productservice/internal/entity/inventory"
)

type inventoryUseCase interface {
	GetInventory(stockID int64) (inventoryentity.Inventory, error)
	UpdateInventory(stockID int64, updatedInventory inventoryentity.Inventory) error
}
type InventoryHandler struct {
	inventoryUseCase inventoryUseCase
}

func New(inventoryUseCase inventoryUseCase) *InventoryHandler {
	return &InventoryHandler{inventoryUseCase}
}

// func (h *InventoryHandler) GetInventory(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	stockID, err := strconv.ParseInt(vars["stockID"], 10, 64)
// 	if err != nil {
// 		http.Error(w, "Invalid stock ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Get the inventory record by stock ID
// 	inventoryRecord, err := h.inventoryUseCase.GetInventory(stockID)
// 	if err != nil {
// 		http.Error(w, "Inventory record not found", http.StatusNotFound)
// 		return
// 	}

// 	// Return the inventory record as a response
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(inventoryRecord)
// }

func (h *InventoryHandler) UpdateInventory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	inventoryID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedInventory inventoryentity.Inventory
	err = json.NewDecoder(r.Body).Decode(&updatedInventory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.inventoryUseCase.UpdateInventory(inventoryID, updatedInventory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
