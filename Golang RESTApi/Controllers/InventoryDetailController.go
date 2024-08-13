package controllers

import "net/http"

type InventoryDetailController interface {
	ChangeStock(w http.ResponseWriter, r *http.Request)
}
