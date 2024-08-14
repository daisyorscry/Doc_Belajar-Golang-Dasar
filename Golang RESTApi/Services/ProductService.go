package services

import (
	requests "RESTApi/Models/Requests"
	responses "RESTApi/Models/Responses"
	"context"
)

type ProductService interface {
	Create(ctx context.Context, request requests.CreateProductRequest) (responses.ProductRespon, error)
	CreateProductWithInventoryDetails(ctx context.Context, request requests.CreateProductRequest) (responses.ProductRespon, error)
	Update(ctx context.Context, request requests.UpdateProductRequest) (responses.ProductRespon, error)
	Delete(ctx context.Context, request int) error
	FindById(ctx context.Context, request int) (responses.ProductRespon, error)
	FindAll(ctx context.Context) ([]responses.ProductRespon, error)
	FindProductDetail(ctx context.Context, request int) (responses.ProductDetailRespon, error)
}
