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
	"github.com/mivanrm/productservice/internal/handler/product"
	"github.com/mivanrm/productservice/internal/handler/review"
	inventoryrepo "github.com/mivanrm/productservice/internal/repo/inventory"
	productrepo "github.com/mivanrm/productservice/internal/repo/product"
	variantrepo "github.com/mivanrm/productservice/internal/repo/variant"

	productuc "github.com/mivanrm/productservice/internal/usecase/product"

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

	productUsecase := productuc.New(&productRepo, &variantRepo, &inventoryRepo)

	productHandler := product.New(&productUsecase)
	r.HandleFunc("/product/{id}", productHandler.GetProduct).Methods("GET")
	r.HandleFunc("/product", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/product", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/product", productHandler.DeleteProduct).Methods("DELETE")

	reviewHandler := review.New()
	r.HandleFunc("/review", reviewHandler.GetReview).Methods("GET")
	r.HandleFunc("/review", reviewHandler.AddReview).Methods("POST")
	r.HandleFunc("/review", reviewHandler.UpdateReview).Methods("PUT")
	r.HandleFunc("/review", reviewHandler.DeleteReview).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
