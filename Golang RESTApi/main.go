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

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
