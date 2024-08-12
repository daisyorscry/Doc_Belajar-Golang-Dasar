package services

import (
	requests "RESTApi/Models/Requests"
	responses "RESTApi/Models/Responses"
	"context"
)

type InventoryProductService interface {
	Create(ctx context.Context, request requests.CreateInventoryProductRequest) (responses.InventoryProductResponse, error)
	FindById(ctx context.Context, id int) (responses.InventoryProductResponse, error)
	FindAll(ctx context.Context) ([]responses.InventoryProductResponse, error)
	Delete(ctx context.Context, id int) error
}
