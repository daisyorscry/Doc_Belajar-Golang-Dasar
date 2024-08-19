package repository

import (
	entity "RESTApi/Models/Entity"
	"context"
	"database/sql"
)

type InventoryProductRepository interface {
	Create(ctx context.Context, tx *sql.Tx, product entity.InventoryProduct) (entity.InventoryProduct, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (entity.InventoryProduct, error)
	FindInventoryByProductId(ctx context.Context, tx *sql.Tx, productId int) (int, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.InventoryProduct, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
}
