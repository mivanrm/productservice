package product

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mivanrm/productservice/internal/entity/product"
)

type productUC interface {
	CreateProduct(product product.CreateProductParam) error
	GetProduct(productID int64) (product.ProductResponse, error)
	UpdateProduct(updatedProduct product.UpdateProductParam) error
	DeleteProduct(productID int64) error
}

type handler struct {
	productUC productUC
}

func New(productUC productUC) *handler {
	return &handler{
		productUC: productUC,
	}
}

func (ph *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product product.CreateProductParam
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := ph.productUC.CreateProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Product created")
}

func (ph *handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	productID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := ph.productUC.GetProduct(productID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (ph *handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updatedProduct product.UpdateProductParam
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ph.productUC.UpdateProduct(updatedProduct); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Product updated successfully")
}

func (ph *handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ph.productUC.DeleteProduct(productID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Product deleted successfully")
}
