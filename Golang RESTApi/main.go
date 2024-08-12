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
	defer db.Close() // Pastikan koneksi database ditutup ketika aplikasi berhenti

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

	// Routes untuk produk
	r.Route("/api/products", func(r chi.Router) {
		r.Use(helper.JWTAuthentication) // Gunakan middleware JWT di route ini
		r.Post("/", productController.Create)
		r.Put("/{id}", productController.Update)
		r.Delete("/{id}", productController.Delete)
		r.Get("/{id}", productController.FindById)
		r.Get("/", productController.FindAll)
	})

	// Routes untuk user
	r.Route("/api/user", func(r chi.Router) {
		r.Use(helper.JWTAuthentication)         // Gunakan middleware JWT di route ini
		r.Put("/update", userController.Update) // Endpoint untuk memperbarui user
	})

	// Routes untuk autentikasi
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/login", userController.Login)               // Endpoint untuk login
		r.Post("/registration", userController.Registration) // Endpoint untuk registrasi
	})

	// Inventory Product routes
	r.Route("/api/inventory-products", func(r chi.Router) {
		r.Post("/", inventoryProductController.Create)
		r.Get("/{id}", inventoryProductController.FindById)
		r.Get("/", inventoryProductController.FindAll)
		// r.Put("/inventory-products", inventoryProductController.Update)
		r.Delete("/{id}", inventoryProductController.Delete)

		// // Inventory Detail routes
		// r.Post("/inventory-details", inventoryDetailController.Create)
		// r.Get("/inventory-details/{id}", inventoryDetailController.FindById)
		// r.Get("/inventory-details/product/{inventory_product_id}", inventoryDetailController.FindAllByProductId)
		// r.Put("/inventory-details", inventoryDetailController.Update)
		// r.Delete("/inventory-details/{id}", inventoryDetailController.Delete)
	})

	r.Post("/api/inventory/stock-change", inventoryDetailController.ChangeStock)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
