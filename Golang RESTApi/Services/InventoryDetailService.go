package services

import (
	requests "RESTApi/Models/Requests"
	responses "RESTApi/Models/Responses"
	"context"
)

type InventoryDetailService interface {
	ChangeStock(ctx context.Context, request requests.StockChangeRequest) error
	FindInventoryDetailById(ctx context.Context, id int) (responses.InventoryDetailResponse, error)

	// CentralizedStockUpdate(ctx context.Context, productId int, change int) error

	// // Create adds a new inventory detail and returns the created detail.
	// Create(ctx context.Context, detail entity.InventoryDetail) (responses.InventoryProductResponse, error)

	// // FindAllByProductId retrieves all inventory details for a specific inventory product.
	// FindAllByProductId(ctx context.Context, inventoryProductId int) ([]responses.InventoryProductResponse, error)

	// // Delete removes an inventory detail by its ID.
	// Delete(ctx context.Context, id int) error
}
