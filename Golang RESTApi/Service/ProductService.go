package service

import (
	requests "RESTApi/Models/Requests"
	responses "RESTApi/Models/Responses"
	"context"
)

type ProductService interface {
	Create(ctx context.Context, request requests.ProductRequest) responses.ProductRespon
	Update(ctx context.Context, request requests.ProductRequest) responses.ProductRespon
	Delete(ctx context.Context, request requests.ProductRequest)
	FindById(ctx context.Context, request requests.ProductRequest) responses.ProductRespon
	FindAll(ctx context.Context) []responses.ProductRespon
}
