package controllers

import "net/http"

type ProductController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindDetailProduct(w http.ResponseWriter, r *http.Request)
	CreateAll(w http.ResponseWriter, r *http.Request)
}
