package repository

import (
	entity "RESTApi/Models/Entity"
	"context"
	"database/sql"
)

// InventoryDetailRepository defines the methods for inventory detail data access.
type InventoryDetailRepository interface {
	FindInventoryByProductId(ctx context.Context, tx *sql.Tx, productId int) (int, error)
	FindByInventoryId(ctx context.Context, tx *sql.Tx, inventoryId int) (entity.InventoryDetail, error)
	UpdateStock(ctx context.Context, tx *sql.Tx, inventoryId int, change int, status string) (int, error)

	// // Create adds a new inventory detail to the database.
	// Create(ctx context.Context, tx *sql.Tx, detail entity.InventoryDetail) (entity.InventoryDetail, error)

	// // FindById retrieves a single inventory detail by its ID.
	// FindById(ctx context.Context, tx *sql.Tx, id int) (entity.InventoryDetail, error)

	// // FindAll retrieves all inventory details for a specific inventory product.
	// FindAllByProductId(ctx context.Context, tx *sql.Tx, inventoryProductId int) ([]entity.InventoryDetail, error)

	// // Update modifies an existing inventory detail in the database.
	// Update(ctx context.Context, tx *sql.Tx, detail entity.InventoryDetail) (entity.InventoryDetail, error)

	// // Delete removes an inventory detail from the database.
	// Delete(ctx context.Context, tx *sql.Tx, id int) error
}
