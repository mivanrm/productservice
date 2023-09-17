package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/mivanrm/productservice/config"
	"github.com/mivanrm/productservice/internal/handler/product"
	productrepo "github.com/mivanrm/productservice/internal/repo/product"
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

	conn, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/product_service", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port))
	if err != nil {
		log.Fatalln(err)
	}

	r := mux.NewRouter()
	productRepo := productrepo.New(conn)
	productUsecase := productuc.New(&productRepo)

	productHandler := product.New(&productUsecase)
	r.HandleFunc("/product", productHandler.GetProduct).Methods("GET")
	r.HandleFunc("/product", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/product", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/product", productHandler.DeleteProduct).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
