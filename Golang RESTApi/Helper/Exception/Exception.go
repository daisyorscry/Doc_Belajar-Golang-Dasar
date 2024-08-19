package exception

import (
	helper "RESTApi/Helper"
	"fmt"
	"log"
	"net/http"
)

type CustomError struct {
	Code    string
	Message string
}

func (error *CustomError) Error() string {
	return fmt.Sprintf("%s: %s", error.Code, error.Message)
}

func NewCustomError(code, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func Exception(w http.ResponseWriter, err error) {
	if customErr, ok := err.(*CustomError); ok {
		switch customErr.Code {
		case "validation_error":
			helper.WriteJsonResponse(w, http.StatusBadRequest, "BAD_REQUEST", customErr.Message)
		case "unauthorized":
			helper.WriteJsonResponse(w, http.StatusUnauthorized, "UNAUTHORIZED", customErr.Message)
		case "database_error":
			helper.WriteJsonResponse(w, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", customErr.Message)
		case "not_found":
			helper.WriteJsonResponse(w, http.StatusNotFound, "NOT_FOUND", customErr.Message)
		case "forbidden":
			helper.WriteJsonResponse(w, http.StatusForbidden, "FORBIDDEN", customErr.Message)
		default:
			helper.WriteJsonResponse(w, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", customErr.Message)
		}
	} else {
		helper.WriteJsonResponse(w, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error())
	}
}

func ServiceErr(err error, message, code string) error {
	log.Printf("%s: %v", message, err)
	return NewCustomError(code, message)
}

func RepositoryErr(err error, context, code string) error {
	log.Printf("%s: %v", context, err)
	return NewCustomError(code, context)
}
