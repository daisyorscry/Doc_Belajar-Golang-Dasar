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

type ProductControllerImpl struct {
	Service services.ProductService
}

func NewProductController(s services.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{Service: s}
}

func (c *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var request requests.CreateProductRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusBadRequest, "fail", "Invalid request payload")
		return
	}

	response, err := c.Service.Create(r.Context(), request)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "success", response)
}

func (c *ProductControllerImpl) CreateAll(w http.ResponseWriter, r *http.Request) {
	var request requests.CreateProductRequest

	// Decode JSON request body
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusBadRequest, "fail", "Invalid request payload")
		return
	}

	// Validasi data jika diperlukan
	// err = validateRequest(request)
	// if err != nil {
	//     helper.WriteJsonResponse(w, http.StatusBadRequest, "fail", err.Error())
	//     return
	// }

	// Call service layer to create product and related details
	response, err := c.Service.CreateProductWithInventoryDetails(r.Context(), request)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	// Write successful response
	helper.WriteJsonResponse(w, http.StatusOK, "success", response)
}

func (c *ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	var request requests.UpdateProductRequest
	request.Id = int(id)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusBadRequest, "fail", "Invalid request payload")
		return
	}

	response, err := c.Service.Update(r.Context(), request)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "success", response)
}

func (c *ProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "Invalid Id", err.Error())
		return
	}

	err = c.Service.Delete(r.Context(), id)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "success", "Product deleted successfully")
}

func (c *ProductControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "Invalid Id", err.Error())
		return
	}

	response, err := c.Service.FindById(r.Context(), id)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "success", response)
}

func (c *ProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	response, err := c.Service.FindAll(r.Context())
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "success", response)
}

func (c *ProductControllerImpl) FindDetailProduct(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "Invalid Id", err.Error())
		return
	}
	response, err := c.Service.FindProductDetail(r.Context(), id)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "error", err.Error())
		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "success", response)
}
