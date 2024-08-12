package repository

import (
	helper "RESTApi/Helper"
	entity "RESTApi/Models/Entity"
	"context"
	"database/sql"
)

type InventoryProductRepositoryImpl struct {
	DB *sql.DB
}

func NewInventoryProductRepository(db *sql.DB) InventoryProductRepository {
	return &InventoryProductRepositoryImpl{
		DB: db,
	}
}

func (r *InventoryProductRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, product entity.InventoryProduct) (entity.InventoryProduct, error) {
	SQL := `
        INSERT INTO inventory_product (product_id, price, created_at, updated_at)
        VALUES ($1, $2, NOW(), NOW())
        RETURNING id, created_at, updated_at
    `
	var inventory entity.InventoryProduct
	err := tx.QueryRowContext(ctx, SQL, product.ProductId, product.Price).Scan(&inventory.Id, &inventory.CreatedAt, &inventory.UpdatedAt)
	if err != nil {
		return entity.InventoryProduct{}, helper.RepositoryErr(err, "error creating inventory product")
	}
	inventory.ProductId = product.ProductId
	inventory.Price = product.Price
	return inventory, nil
}

func (r *InventoryProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (entity.InventoryProduct, error) {
	SQL := `
		SELECT id, product_id, price, created_at, updated_at
		FROM inventory_product
		WHERE id = $1
		`
	var inventory entity.InventoryProduct
	err := tx.QueryRowContext(ctx, SQL, id).Scan(&inventory.Id, &inventory.ProductId, &inventory.Price, &inventory.CreatedAt, &inventory.UpdatedAt)
	if err != nil {
		return inventory, helper.RepositoryErr(err, "error finding product inventory by id")
	}
	return inventory, nil
}

func (r *InventoryProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.InventoryProduct, error) {
	SQL := `
		SELECT id, product_id, price, created_at, updated_at
		FROM inventory_product
		`
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, helper.RepositoryErr(err, "error finding all inventory products")
	}
	defer rows.Close()

	var products []entity.InventoryProduct
	for rows.Next() {
		var product entity.InventoryProduct
		err := rows.Scan(&product.Id, &product.ProductId, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, helper.RepositoryErr(err, "error scanning inventory products")
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *InventoryProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	SQL := `
		DELETE FROM inventory_product
		WHERE id = $1
		`
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return helper.RepositoryErr(err, "error deleting inventory product")
	}
	return nil
}
