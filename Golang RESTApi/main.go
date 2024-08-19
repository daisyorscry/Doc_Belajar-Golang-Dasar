package main

import (
	auth "RESTApi/Auth"
	config "RESTApi/Config"
	controllers "RESTApi/Controllers"
	repository "RESTApi/Models/Repository"
	services "RESTApi/Services"
	"database/sql"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func main() {

	// initialize chii router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// initial posgres database
	db, err := sql.Open("postgres", "user=postgres password=123 dbname=belajar_golang sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(1000)
	db.SetConnMaxIdleTime(time.Second * 5)
	db.SetConnMaxLifetime(time.Minute * 60)

	// initialize redis database
	redisClient := config.NewRedisClient()

	// initialize validator
	validate := validator.New()

	// Initialize repositories
	userRepository := repository.NewUserRepository()
	productRepository := repository.NewProductRepository()
	inventoryProductRepo := repository.NewInventoryProductRepository()
	inventoryDetailRepo := repository.NewInventoryDetailRepository()

	// Initialize services
	userService := services.NewUserService(userRepository, db, validate)
	productService := services.NewProductService(productRepository, inventoryProductRepo, inventoryDetailRepo, userRepository, db, validate)
	inventoryProductService := services.NewInventoryProductService(inventoryProductRepo, db)
	inventoryDetailService := services.NewInventoryDetailService(inventoryDetailRepo, inventoryProductRepo, db, redisClient)

	// Initialize controllers
	userController := controllers.NewUserController(userService)
	productController := controllers.NewProductController(productService)
	inventoryProductController := controllers.NewInventoryProductController(inventoryProductService)
	inventoryDetailController := controllers.NewInventoryDetailController(inventoryDetailService)

	// Routes untuk user
	r.Route("/api/user", func(r chi.Router) {
		r.Use(auth.Auth)
		r.Put("/update", userController.Update)
	})

	// Routes untuk autentikasi
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/login", userController.Login)
		r.Post("/registration", userController.Registration)
	})

	r.Route("/api/products", func(r chi.Router) {
		r.Use(auth.Auth)
		r.Post("/", productController.Create)
		r.Post("/details", productController.CreateAll)
		r.Put("/{id}", productController.Update)
		r.Delete("/{id}", productController.Delete)
		r.Get("/{id}", productController.FindById)
		r.Get("/detail/{id}", productController.FindDetailProduct)
		r.Get("/", productController.FindAll)
	})

	// Inventory Product routes
	r.Route("/api/inventory-products", func(r chi.Router) {
		r.Use(auth.Auth)
		r.Post("/", inventoryProductController.Create)
		r.Get("/{id}", inventoryProductController.FindById)
		r.Get("/", inventoryProductController.FindAll)
		r.Delete("/{id}", inventoryProductController.Delete)
	})

	// Inventory Details routes
	r.Route("/api/inventory-details", func(r chi.Router) {
		r.Use(auth.Auth)
		r.Get("/{id}", inventoryDetailController.FindInventoryDetailById)
		r.Post("/stock-change", inventoryDetailController.ChangeStock)

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
