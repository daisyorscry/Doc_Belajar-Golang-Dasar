package main

import (
	controllers "RESTApi/Controllers"
	repository "RESTApi/Models/Repository"
	services "RESTApi/Services"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq" // Driver untuk PostgreSQL
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db, err := sql.Open("postgres", "user=postgres password=123 dbname=belajar_golang sslmode=disable")
	if err != nil {
		panic(err)
	}

	Service := &services.ProductServiceImpl{
		ProductRepository: &repository.ProductRepositoryImpl{},
		DB:                db,
		Validate:          validator.New(),
	}

	productController := &controllers.ProductControllerImpl{
		Service: Service,
	}

	r.Route("/api/products", func(r chi.Router) {
		r.Post("/", productController.Create)
		r.Put("/{id}", productController.Update)
		r.Delete("/{id}", productController.Delete)
		r.Get("/{id}", productController.FindById)
		r.Get("/", productController.FindAll)
	})
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
