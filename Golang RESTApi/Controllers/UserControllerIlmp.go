package controllers

import (
	helper "RESTApi/Helper"
	requests "RESTApi/Models/Requests"
	services "RESTApi/Services"
	"encoding/json"
	"net/http"
)

type UserControllerImpl struct {
	Service services.UserService
}

func (c *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var request requests.UserLoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusBadRequest, "error", "Invalid request payload")
		return
	}

	response, token, err := c.Service.Login(r.Context(), request)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	w.Header().Set("X-API-TOKEN", "Bearer "+token)
	helper.WriteJsonResponse(w, http.StatusOK, "success", response)
}

func (c *UserControllerImpl) Registration(w http.ResponseWriter, r *http.Request) {
	var request requests.UserRegistrationRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusBadRequest, "error", "Invalid request payload")
		return
	}

	responses, err := c.Service.Register(r.Context(), request)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "error registration", err.Error())
		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "sukses", responses)

}

func (c *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	var request requests.UserUpdateRequest

	// Decode JSON request body
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusBadRequest, "error", "Invalid request payload")
		return
	}

	// Panggil service untuk melakukan update
	response, err := c.Service.Update(r.Context(), request)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	// Kirimkan respon sukses
	helper.WriteJsonResponse(w, http.StatusOK, "success", response)
}
