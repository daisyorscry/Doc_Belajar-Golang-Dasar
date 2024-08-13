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

type InventoryDetailControllerImpl struct {
	Service services.InventoryDetailService
}

func NewInventoryDetailController(s services.InventoryDetailService) *InventoryDetailControllerImpl {
	return &InventoryDetailControllerImpl{Service: s}
}

// ******************************CONTROLLER INVENTORY DETAILS********************************************

// ******************************CHANGE STOCK CONTROLLER********************************************

func (c *InventoryDetailControllerImpl) ChangeStock(w http.ResponseWriter, r *http.Request) {
	var request requests.StockChangeRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = c.Service.ChangeStock(r.Context(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.WriteJsonResponse(w, http.StatusOK, "OK", "stock upadte sukses")
}

// ******************************FIND INVENTORY DETAIL BY ID CONTROLLER********************************************

func (c *InventoryDetailControllerImpl) FindInventoryDetailById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusBadRequest, "BAD REQUEST", "invalid  request payload")
		return
	}

	product, err := c.Service.FindInventoryDetailById(r.Context(), id)
	if err != nil {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "something wring", err.Error())

		return
	}
	helper.WriteJsonResponse(w, http.StatusOK, "OK", product)

}
