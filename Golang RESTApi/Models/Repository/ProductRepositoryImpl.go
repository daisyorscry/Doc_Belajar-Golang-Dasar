package repository

import (
	helper "RESTApi/Helper"
	entity "RESTApi/Models/Entity"
	"context"
	"database/sql"
	"fmt"
)

type ProductRepositoryImpl struct{}

func (r *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product entity.Product) (entity.Product, error) {
	SQL := "INSERT INTO product (ProductName, ProductDesc) value (?, ?) RETURNING id"

	var id int64
	err := tx.QueryRowContext(ctx, SQL, product.ProductName, product.ProductDesc).Scan(&id)
	if err != nil {
		return entity.Product{}, err
	}

	product.Id = id
	return product, nil
}

func (r *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product entity.Product) (entity.Product, error) {
	SQL := "UPDATE product SET ProductName = ?, ProductDesc = ? WHERE id = ? RETURNING id"

	var id int

	err := tx.QueryRowContext(ctx, SQL, product.ProductName, product.ProductDesc, product.Id).Scan(&id)
	if err != nil {
		return entity.Product{}, err
	}

	product.Id = int64(id)
	return product, nil
}

func (r *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product entity.Product) {
	SQL := "DELETE FROM product WHERE id = ?"

	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.DatabaseErr(err)

}
func (r *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (entity.Product, error) {
	SQL := "SELECT id, ProductName, ProductDesc FROM product WHERE id = ?"

	product := entity.Product{}
	err := tx.QueryRowContext(ctx, SQL, productId).Scan(&product.Id, &product.ProductName, &product.ProductDesc)
	if err != nil {
		if err == sql.ErrNoRows {
			return product, fmt.Errorf("product with ID %d not found", productId)
		}
		return product, helper.DatabaseErr(err)
	}

	return product, nil

}
func (r *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Product {
	SQL := "SELECT id, ProductName,ProductDesc"

	result, err := tx.QueryContext(ctx, SQL)
	helper.DatabaseErr(err)

	var products []entity.Product
	for result.Next() {
		product := entity.Product{}
		err := result.Scan(&product.Id, product.ProductName, product.ProductDesc)
		helper.DatabaseErr(err)
		products = append(products, product)
	}
	return products
}
