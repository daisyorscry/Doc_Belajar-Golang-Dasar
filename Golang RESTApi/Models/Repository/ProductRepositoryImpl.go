package repository

import (
	helper "RESTApi/Helper"
	entity "RESTApi/Models/Entity"
	"context"
	"database/sql"
)

type ProductRepositoryImpl struct{}

func (r *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product entity.Product) (entity.Product, error) {
	SQL := "INSERT INTO product (productname, productdesc) VALUES ($1, $2) RETURNING id"

	var id int64
	err := tx.QueryRowContext(ctx, SQL, product.ProductName, product.ProductDesc).Scan(&id)
	if err != nil {
		return entity.Product{}, helper.RepositoryErr(err, "error create product")
	}

	product.Id = id
	return product, nil
}

func (r *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product entity.Product) (entity.Product, error) {
	SQL := "UPDATE product SET productname = $1, productdesc = $2 WHERE id = $3 RETURNING id"

	var id int
	err := tx.QueryRowContext(ctx, SQL, product.ProductName, product.ProductDesc, product.Id).Scan(&id)
	if err != nil {
		return entity.Product{}, helper.RepositoryErr(err, "error updating product")
	}

	product.Id = int64(id)
	return product, nil
}

func (r *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product entity.Product) error {
	SQL := "DELETE FROM product WHERE id = $1"

	_, err := tx.ExecContext(ctx, SQL, product.Id)
	if err != nil {
		return helper.RepositoryErr(err, "deleting product")
	}
	return nil
}

func (r *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (entity.Product, error) {
	SQL := "SELECT * FROM product WHERE id = $1"

	product := entity.Product{}
	err := tx.QueryRowContext(ctx, SQL, productId).Scan(&product.Id, &product.ProductName, &product.ProductDesc)
	if err != nil {
		return entity.Product{}, helper.RepositoryErr(err, "error finding product by id")
	}

	return product, nil
}

func (r *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Product, error) {
	SQL := "SELECT id, ProductName,ProductDesc from product"

	result, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, helper.RepositoryErr(err, "error finding all products")
	}
	defer result.Close()

	var products []entity.Product
	for result.Next() {
		product := entity.Product{}
		err := result.Scan(&product.Id, &product.ProductName, &product.ProductDesc)
		if err != nil {
			return nil, helper.RepositoryErr(err, "error finding all products")
		}
		products = append(products, product)
	}
	return products, nil
}
