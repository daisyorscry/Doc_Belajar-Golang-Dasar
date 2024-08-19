package controllers

import (
	helper "RESTApi/Helper"
	exception "RESTApi/Helper/Exception"
	requests "RESTApi/Models/Requests"
	services "RESTApi/Services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductControllerImpl struct {
	Service services.ProductService
}

func NewProductController(s services.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{Service: s}
}

// controlller for create product
func (c *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var request requests.CreateProductRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		exception.Exception(w, err)
		return
	}

	response, err := c.Service.Create(r.Context(), request)
	if err != nil {
		exception.Exception(w, err)
		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "OK", response)
}

// controller for create product with inventory and stock
func (c *ProductControllerImpl) CreateAll(w http.ResponseWriter, r *http.Request) {
	var request requests.CreateProductRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		exception.Exception(w, err)

		return
	}

	response, err := c.Service.CreateProductWithInventoryDetails(r.Context(), request)
	if err != nil {
		exception.Exception(w, err)

		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "OK", response)
}

// controller product for update
func (c *ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		exception.Exception(w, err)

		return
	}
	var request requests.UpdateProductRequest
	request.Id = int(id)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		exception.Exception(w, err)

		return
	}

	response, err := c.Service.Update(r.Context(), request)
	if err != nil {
		exception.Exception(w, err)

		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "OK", response)
}

// controller product for delete
func (c *ProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		exception.Exception(w, err)

		return
	}

	err = c.Service.Delete(r.Context(), id)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "OK", "Product deleted successfully")
}

// controller product for find by id
func (c *ProductControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		exception.Exception(w, err)

		return
	}

	response, err := c.Service.FindById(r.Context(), id)
	if err != nil {
		exception.Exception(w, err)

		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "OK", response)
}

// controller product for get all
func (c *ProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	response, err := c.Service.FindAll(r.Context())
	if err != nil {
		exception.Exception(w, err)

		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "OK", response)
}

// controller product for get all product details
func (c *ProductControllerImpl) FindDetailProduct(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		exception.Exception(w, err)

		return
	}
	response, err := c.Service.FindProductDetail(r.Context(), id)
	if err != nil {
		exception.Exception(w, err)

		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "OK", response)
}
