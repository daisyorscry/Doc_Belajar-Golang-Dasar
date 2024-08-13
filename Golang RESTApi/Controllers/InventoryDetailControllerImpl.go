package controllers

import (
	helper "RESTApi/Helper"
	requests "RESTApi/Models/Requests"
	services "RESTApi/Services"
	"encoding/json"
	"net/http"
)

type InventoryDetailControllerImpl struct {
	Service services.InventoryDetailService
}

func NewInventoryDetailController(s services.InventoryDetailService) *InventoryDetailControllerImpl {
	return &InventoryDetailControllerImpl{Service: s}
}

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
