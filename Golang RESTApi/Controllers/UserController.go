package controllers

import "net/http"

type UserController interface {
	Login(w http.ResponseWriter, r *http.Request)
	Registration(w http.ResponseWriter, r *http.Request)

	// Delete(w http.ResponseWriter, r *http.Request)
	// Update(w http.ResponseWriter, r *http.Request)
	// FindById(w http.ResponseWriter, r *http.Request)
	// FindAll(w http.ResponseWriter, r *http.Request)
}
