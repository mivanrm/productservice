package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mivanrm/productservice/config"
	inventoryhandler "github.com/mivanrm/productservice/internal/handler/inventory"
	"github.com/mivanrm/productservice/internal/handler/product"
	reviewhandler "github.com/mivanrm/productservice/internal/handler/review"

	inventoryrepo "github.com/mivanrm/productservice/internal/repo/inventory"
	productrepo "github.com/mivanrm/productservice/internal/repo/product"
	reviewrepo "github.com/mivanrm/productservice/internal/repo/review"
	variantrepo "github.com/mivanrm/productservice/internal/repo/variant"

	inventoryuc "github.com/mivanrm/productservice/internal/usecase/inventory"
	productuc "github.com/mivanrm/productservice/internal/usecase/product"
	reviewuc "github.com/mivanrm/productservice/internal/usecase/review"

	"gopkg.in/yaml.v2"
)

func main() {

	f, err := os.Open("files/config.yml")
	if err != nil {
		log.Fatalln(err)

	}
	defer f.Close()

	var cfg config.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := sqlx.Connect("postgres", fmt.Sprintf("postgresql://%s:%s@%s:%s/product_service?sslmode=disable", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port))
	if err != nil {
		log.Fatalln(err)
	}

	r := mux.NewRouter()
	productRepo := productrepo.New(conn)
	variantRepo := variantrepo.New(conn)
	inventoryRepo := inventoryrepo.New(conn)
	reviewRepo := reviewrepo.New(conn)

	productUsecase := productuc.New(&productRepo, &variantRepo, &inventoryRepo)
	inventoryUsecase := inventoryuc.New(&inventoryRepo)
	reviewUsecase := reviewuc.New(&reviewRepo)

	productHandler := product.New(&productUsecase)

	r.HandleFunc("/product/{id}", productHandler.GetProduct).Methods("GET")
	r.HandleFunc("/product", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/product", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/product/{id}", productHandler.DeleteProduct).Methods("DELETE")

	reviewHandler := reviewhandler.New(&reviewUsecase)
	r.HandleFunc("/review", reviewHandler.GetReview).Methods("GET")
	r.HandleFunc("/review", reviewHandler.CreateReview).Methods("POST")
	r.HandleFunc("/review", reviewHandler.UpdateReview).Methods("PUT")
	r.HandleFunc("/review", reviewHandler.DeleteReview).Methods("DELETE")

	inventoryHandler := inventoryhandler.New(&inventoryUsecase)
	r.HandleFunc("/inventory", inventoryHandler.UpdateInventory).Methods("PUT")

	http.Handle("/", r)
	fmt.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
