package services

import (
	requests "RESTApi/Models/Requests"
	"context"
)

type InventoryDetailService interface {
	ChangeStock(ctx context.Context, request requests.StockChangeRequest) error

	// // Create adds a new inventory detail and returns the created detail.
	// Create(ctx context.Context, detail entity.InventoryDetail) (responses.InventoryProductResponse, error)

	// // FindById retrieves a single inventory detail by its ID.
	// FindById(ctx context.Context, id int) (entity.InventoryDetail, error)

	// // FindAllByProductId retrieves all inventory details for a specific inventory product.
	// FindAllByProductId(ctx context.Context, inventoryProductId int) ([]responses.InventoryProductResponse, error)

	// // Update modifies an existing inventory detail and returns the updated detail.
	// Update(ctx context.Context, detail entity.InventoryDetail) (responses.InventoryProductResponse, error)

	// // Delete removes an inventory detail by its ID.
	// Delete(ctx context.Context, id int) error
}
