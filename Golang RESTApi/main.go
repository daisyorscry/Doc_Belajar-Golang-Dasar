package main

import (
	controllers "RESTApi/Controllers"
	helper "RESTApi/Helper"
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
	defer db.Close()

	ProductService := &services.ProductServiceImpl{
		ProductRepository: &repository.ProductRepositoryImpl{},
		DB:                db,
		Validate:          validator.New(),
	}

	productController := &controllers.ProductControllerImpl{
		Service: ProductService,
	}

	UserService := &services.UserServiceImpl{
		UserRepository: &repository.UserRepositoryImpl{},
		DB:             db,
		Validate:       validator.New(),
	}

	userController := &controllers.UserControllerImpl{
		Service: UserService,
	}

	// Initialize repositories
	inventoryProductRepo := repository.NewInventoryProductRepository(db)
	inventoryDetailRepo := repository.NewInventoryDetailRepository(db)

	// Initialize services
	inventoryProductService := services.NewInventoryProductService(inventoryProductRepo, db)
	inventoryDetailService := services.NewInventoryDetailService(inventoryDetailRepo, db)

	// Initialize controllers
	inventoryProductController := controllers.NewInventoryProductController(inventoryProductService)
	inventoryDetailController := controllers.NewInventoryDetailController(inventoryDetailService)

	// Routes untuk user
	r.Route("/api/user", func(r chi.Router) {
		r.Use(helper.JWTAuthentication)
		r.Put("/update", userController.Update)
	})

	// Routes untuk autentikasi
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/login", userController.Login)
		r.Post("/registration", userController.Registration)
	})

	r.Route("/api/products", func(r chi.Router) {
		r.Use(helper.JWTAuthentication)
		r.Post("/", productController.Create)
		r.Put("/{id}", productController.Update)
		r.Delete("/{id}", productController.Delete)
		r.Get("/{id}", productController.FindById)
		r.Get("/", productController.FindAll)
	})

	// Inventory Product routes
	r.Route("/api/inventory-products", func(r chi.Router) {
		r.Post("/", inventoryProductController.Create)
		r.Get("/{id}", inventoryProductController.FindById)
		r.Get("/", inventoryProductController.FindAll)
		r.Delete("/{id}", inventoryProductController.Delete)
	})

	// Inventory Details routes
	r.Route("/api/inventory-details", func(r chi.Router) {
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
