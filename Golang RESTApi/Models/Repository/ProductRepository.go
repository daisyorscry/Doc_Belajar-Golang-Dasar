package repository

import (
	entity "RESTApi/Models/Entity"
	"context"
	"database/sql"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product entity.Product) (entity.Product, error)
	Update(ctx context.Context, tx *sql.Tx, product entity.Product) (entity.Product, error)
	Delete(ctx context.Context, tx *sql.Tx, product entity.Product)
	FindById(ctx context.Context, tx *sql.Tx, productId int) (entity.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Product
}
