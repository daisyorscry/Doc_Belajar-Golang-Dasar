package controllers

import (
	helper "RESTApi/Helper"
	requests "RESTApi/Models/Requests"
	services "RESTApi/Services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type InventoryProductControllerImpl struct {
	Service services.InventoryProductService
}

func NewInventoryProductController(s services.InventoryProductService) *InventoryProductControllerImpl {
	return &InventoryProductControllerImpl{Service: s}
}

func (c *InventoryProductControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var request requests.CreateInventoryProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helper.WriteJsonResponse(w, http.StatusBadRequest, "BAD REQUEST", "invalid request payload")
		return
	}

	newProduct, err := c.Service.Create(r.Context(), request)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "something wrong", err.Error())
		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "OK", newProduct)
}

func (c *InventoryProductControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusBadRequest, "BAD REQUEST", "invalid  request payload")
		return
	}

	product, err := c.Service.FindById(r.Context(), id)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "something wring", err.Error())

		return
	}
	helper.WriteJsonResponse(w, http.StatusOK, "OK", product)

}

func (c *InventoryProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	products, err := c.Service.FindAll(r.Context())
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusBadRequest, "BAD REQUEST", err.Error())
		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "OK", products)

}

func (c *InventoryProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "something wring", err.Error())
		return
	}

	if err := c.Service.Delete(r.Context(), id); err != nil {
		helper.WriteJsonResponse(w, http.StatusBadRequest, "BAD REQUEST", err.Error())
		return
	}

}
