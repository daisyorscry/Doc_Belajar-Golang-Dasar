package repository

import (
	exception "RESTApi/Helper/Exception"
	entity "RESTApi/Models/Entity"
	"context"
	"database/sql"
	"log"
)

type ProductRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (r *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product entity.Product, userId int) (entity.Product, error) {
	SQL := "INSERT INTO product (productname, productdesc, user_id) VALUES ($1, $2, $3) RETURNING id"

	var id int64
	err := tx.QueryRowContext(ctx, SQL, product.ProductName, product.ProductDesc, userId).Scan(&id)
	if err != nil {
		return entity.Product{}, exception.RepositoryErr(err, "failed create product", "database_error")
	}

	log.Println(product.UserId)

	product.Id = id
	return product, nil
}

func (r *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product entity.Product) (entity.Product, error) {
	SQL := "UPDATE product SET productname = $1, productdesc = $2 WHERE id = $3 RETURNING id"

	var id int
	err := tx.QueryRowContext(ctx, SQL, product.ProductName, product.ProductDesc, product.Id).Scan(&id)
	if err != nil {
		return entity.Product{}, exception.RepositoryErr(err, "failed updating product", "database_error")
	}

	product.Id = int64(id)
	return product, nil
}

func (r *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product entity.Product) error {
	SQL := "DELETE FROM product WHERE id = $1"

	_, err := tx.ExecContext(ctx, SQL, product.Id)
	if err != nil {
		return exception.RepositoryErr(err, "failed delete product", "database_error")
	}
	return nil
}

func (r *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (entity.Product, error) {
	SQL := `
        SELECT 
			p.id, 
			p.productname, 
			p.productdesc, 
			p.created_at, 
			u.username 
        FROM 
			product p
        JOIN 
			users u 
		ON 
			p.user_id = u.id
        WHERE 
			p.id = $1
    `

	product := entity.Product{}
	err := tx.QueryRowContext(ctx, SQL, productId).Scan(
		&product.Id,
		&product.ProductName,
		&product.ProductDesc,
		&product.CreatedAt,
		&product.CreateBy,
	)
	if err != nil {
		return entity.Product{}, exception.RepositoryErr(err, "product not found", "not_found")
	}

	return product, nil
}

func (r *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Product, error) {
	SQL := `
		SELECT 
			p.id, 
			p.ProductName, 
			p.ProductDesc, 
			u.username,
			p.created_at
		FROM 
			product p 
		JOIN 
			users u 
		ON 
			p.user_id = u.id
		`

	result, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, exception.RepositoryErr(err, "failed get product", "database_error")
	}
	defer result.Close()

	var products []entity.Product
	for result.Next() {
		product := entity.Product{}

		err := result.Scan(
			&product.Id,
			&product.ProductName,
			&product.ProductDesc,
			&product.CreateBy,
			&product.CreatedAt,
		)
		if err != nil {
			return nil, exception.RepositoryErr(err, "failed get product", "database_error")
		}

		products = append(products, product)
	}
	return products, nil
}
