package controllers

import (
	helper "RESTApi/Helper"
	exception "RESTApi/Helper/Exception"
	requests "RESTApi/Models/Requests"
	services "RESTApi/Services"
	"encoding/json"
	"net/http"
)

type UserControllerImpl struct {
	Service services.UserService
}

func NewUserController(s services.UserService) *UserControllerImpl {
	return &UserControllerImpl{Service: s}
}

func (c *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var request requests.UserLoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		exception.Exception(w, err)

		return
	}

	response, token, err := c.Service.Login(r.Context(), request)
	if err != nil {
		exception.Exception(w, err)
		return
	}

	w.Header().Set("X-API-TOKEN", "Bearer "+token)
	helper.WriteJsonResponse(w, http.StatusOK, "OK", response)
}

func (c *UserControllerImpl) Registration(w http.ResponseWriter, r *http.Request) {
	var request requests.UserRegistrationRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		exception.Exception(w, err)

		return
	}

	responses, err := c.Service.Register(r.Context(), request)
	if err != nil {
		exception.Exception(w, err)

		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "OK", responses)

}

func (c *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	var request requests.UserUpdateRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		exception.Exception(w, err)

		return
	}

	response, err := c.Service.Update(r.Context(), request)
	if err != nil {
		exception.Exception(w, err)

	}

	helper.WriteJsonResponse(w, http.StatusOK, "OK", response)
}
