package repository

import (
	helper "RESTApi/Helper"
	entity "RESTApi/Models/Entity"
	"context"
	"database/sql"
)

type InventoryDetailRepositoryImpl struct {
	DB *sql.DB
}

func NewInventoryDetailRepository(db *sql.DB) InventoryDetailRepository {
	return &InventoryDetailRepositoryImpl{
		DB: db,
	}
}

func (r *InventoryDetailRepositoryImpl) FindInventoryByProductId(ctx context.Context, tx *sql.Tx, productId int) (int, error) {
	SQL := "SELECT id FROM inventory_product WHERE product_id = $1 FOR UPDATE NOWAIT"
	var inventoryId int
	err := tx.QueryRowContext(ctx, SQL, productId).Scan(&inventoryId)
	if err != nil {
		return 0, helper.RepositoryErr(err, "error fetching inventory by product_id")
	}
	return inventoryId, nil
}

func (r *InventoryDetailRepositoryImpl) FindByInventoryId(ctx context.Context, tx *sql.Tx, inventoryId int) (entity.InventoryDetail, error) {
	SQL := `
		SELECT id, inventory_id, stock, status, created_at, updated_at
		FROM inventory_details
		WHERE inventory_id = $1
		FOR UPDATE NOWAIT
		`

	var inventoryDetail entity.InventoryDetail
	err := tx.QueryRowContext(ctx, SQL, inventoryId).Scan(
		&inventoryDetail.Id,
		&inventoryDetail.InventoryProductId,
		&inventoryDetail.Stock,
		&inventoryDetail.Status,
		&inventoryDetail.CreatedAt,
		&inventoryDetail.UpdatedAt,
	)
	if err != nil {
		return entity.InventoryDetail{}, helper.RepositoryErr(err, "error get inventory_detail by inventory_id")
	}
	return inventoryDetail, nil
}

func (r *InventoryDetailRepositoryImpl) UpdateStock(ctx context.Context, tx *sql.Tx, detail entity.InventoryDetail) (entity.InventoryDetail, error) {
	SQL := `
        UPDATE inventory_details
        SET stock = $1, status = $2, updated_at = NOW()
        WHERE inventory_id = $3
        RETURNING id
    `

	err := tx.QueryRowContext(ctx, SQL, detail.Stock, detail.Status, detail.InventoryProductId).Scan(&detail.Id)
	if err != nil {
		return entity.InventoryDetail{}, helper.RepositoryErr(err, "error updating stock and status")
	}

	return detail, nil
}

// func (r *InventoryDetailRepositoryImpl) FetchCurrentStock(ctx context.Context, tx *sql.Tx, detail entity.InventoryDetail) (entity.InventoryDetail, error) {
// 	SQL := `
//         SELECT stock, status
//         FROM inventory_details
//         WHERE inventory_id = $1
// 		RETRUNING id
//     `
// 	var currentStock int
// 	var status string

// 	err := tx.QueryRowContext(ctx, SQL, inventoryId, status).Scan(&currentStock)
// 	if err != nil {
// 		return 0, helper.RepositoryErr(err, "error fetching current stock")
// 	}
// 	return currentStock, nil
// }
// func (r *InventoryDetailRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, detail entity.InventoryDetail) (entity.InventoryDetail, error) {
// SQL := `
// INSERT INTO inventory_detail (inventory_product_id, quantity, status, created_at, updated_at)
// VALUES ($1, $2, $3, NOW(), NOW())
// RETURNING id
// `
// var id int
// err := tx.QueryRowContext(ctx, SQL, detail.InventoryProductId, detail.Quantity, detail.Status).Scan(&id)
// if err != nil {
// return entity.InventoryDetail{}, helper.RepositoryErr(err, "error creating inventory detail")
// }
// detail.Id = id
// return detail, nil
// }
//
// func (r *InventoryDetailRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (entity.InventoryDetail, error) {
// SQL := `
// SELECT id, inventory_product_id, quantity, status, created_at, updated_at
// FROM inventory_detail
// WHERE id = $1
// `
// var detail entity.InventoryDetail
// err := tx.QueryRowContext(ctx, SQL, id).Scan(&detail.Id, &detail.InventoryProductId, &detail.Quantity, &detail.Status, &detail.CreatedAt, &detail.UpdatedAt)
// if err != nil {
// return detail, helper.RepositoryErr(err, "error finding inventory detail by id")
// }
// return detail, nil
// }
//
// func (r *InventoryDetailRepositoryImpl) FindAllByProductId(ctx context.Context, tx *sql.Tx, inventoryProductId int) ([]entity.InventoryDetail, error) {
// SQL := `
// SELECT id, inventory_product_id, quantity, status, created_at, updated_at
// FROM inventory_detail
// WHERE inventory_product_id = $1
// `
// rows, err := tx.QueryContext(ctx, SQL, inventoryProductId)
// if err != nil {
// return nil, helper.RepositoryErr(err, "error finding inventory details by product id")
// }
// defer rows.Close()
//
// var details []entity.InventoryDetail
// for rows.Next() {
// var detail entity.InventoryDetail
// err := rows.Scan(&detail.Id, &detail.InventoryProductId, &detail.Quantity, &detail.Status, &detail.CreatedAt, &detail.UpdatedAt)
// if err != nil {
// return nil, helper.RepositoryErr(err, "error scanning inventory details")
// }
// details = append(details, detail)
// }
// return details, nil
// }
//
// func (r *InventoryDetailRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, detail entity.InventoryDetail) (entity.InventoryDetail, error) {
// SQL := `
// UPDATE inventory_detail
// SET quantity = $1, status = $2, updated_at = NOW()
// WHERE id = $3
// `
// _, err := tx.ExecContext(ctx, SQL, detail.Quantity, detail.Status, detail.Id)
// if err != nil {
// return detail, helper.RepositoryErr(err, "error updating inventory detail")
// }
// return detail, nil
// }
//
// func (r *InventoryDetailRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
// SQL := `
// DELETE FROM inventory_detail
// WHERE id = $1
// `
// _, err := tx.ExecContext(ctx, SQL, id)
// if err != nil {
// return helper.RepositoryErr(err, "error deleting inventory detail")
// }
// return nil
// }
//
