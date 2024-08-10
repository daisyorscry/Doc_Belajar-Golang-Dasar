package helper

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func WriteJsonResponse(w http.ResponseWriter, code int, status string, data interface{}) {
	response := JsonResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
