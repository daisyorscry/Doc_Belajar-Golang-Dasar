package services

import (
	helper "RESTApi/Helper"
	repository "RESTApi/Models/Repository"
	requests "RESTApi/Models/Requests"
	"context"
	"database/sql"
)

type InventoryDetailServiceImpl struct {
	Repo repository.InventoryDetailRepository
	DB   *sql.DB
}

func NewInventoryDetailService(repo repository.InventoryDetailRepository, db *sql.DB) InventoryDetailService {
	return &InventoryDetailServiceImpl{
		Repo: repo,
		DB:   db,
	}
}

func (s *InventoryDetailServiceImpl) ChangeStock(ctx context.Context, request requests.StockChangeRequest) error {

	tx, err := s.DB.BeginTx(ctx, helper.BeginTxHandlerExec())

	if err != nil {
		return helper.ServiceErr(err, "error starting transaction")
	}
	defer helper.TxHandler(tx, err)

	inventoryId, err := s.Repo.FindInventoryByProductId(ctx, tx, request.ProductId)
	if err != nil {
		return helper.ServiceErr(err, "error fetching inventory by product_id")
	}

	inventoryDetail, err := s.Repo.FindByInventoryId(ctx, tx, inventoryId)
	if err != nil {
		return helper.ServiceErr(err, "error fetching inventory details by inventory_id")
	}

	_, err = s.Repo.UpdateStock(ctx, tx, inventoryId, request.Change, inventoryDetail.Status)
	if err != nil {
		return helper.ServiceErr(err, err.Error())
	}

	return nil
}

// func (s *InventoryDetailServiceImpl) Create(ctx context.Context, detail entity.InventoryDetail) (entity.InventoryDetail, error) {
// 	tx, err := s.DB.BeginTx(ctx, nil)
// 	if err != nil {
// 		return entity.InventoryDetail{}, err
// 	}
// 	defer tx.Rollback()

// 	newDetail, err := s.Repo.Create(ctx, tx, detail)
// 	if err != nil {
// 		return entity.InventoryDetail{}, err
// 	}

// 	if err := tx.Commit(); err != nil {
// 		return entity.InventoryDetail{}, err
// 	}
// 	return newDetail, nil
// }

// func (s *InventoryDetailServiceImpl) FindById(ctx context.Context, id int) (entity.InventoryDetail, error) {
// 	tx, err := s.DB.BeginTx(ctx, nil)
// 	if err != nil {
// 		return entity.InventoryDetail{}, err
// 	}
// 	defer tx.Rollback()

// 	detail, err := s.Repo.FindById(ctx, tx, id)
// 	if err != nil {
// 		return entity.InventoryDetail{}, err
// 	}

// 	return detail, nil
// }

// func (s *InventoryDetailServiceImpl) FindAllByProductId(ctx context.Context, inventoryProductId int) ([]entity.InventoryDetail, error) {
// 	tx, err := s.DB.BeginTx(ctx, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer tx.Rollback()

// 	details, err := s.Repo.FindAllByProductId(ctx, tx, inventoryProductId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return details, nil
// }

// func (s *InventoryDetailServiceImpl) Update(ctx context.Context, detail entity.InventoryDetail) (entity.InventoryDetail, error) {
// 	tx, err := s.DB.BeginTx(ctx, nil)
// 	if err != nil {
// 		return entity.InventoryDetail{}, err
// 	}
// 	defer tx.Rollback()

// 	updatedDetail, err := s.Repo.Update(ctx, tx, detail)
// 	if err != nil {
// 		return entity.InventoryDetail{}, err
// 	}

// 	if err := tx.Commit(); err != nil {
// 		return entity.InventoryDetail{}, err
// 	}
// 	return updatedDetail, nil
// }

// func (s *InventoryDetailServiceImpl) Delete(ctx context.Context, id int) error {
// 	tx, err := s.DB.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()

// 	if err := s.Repo.Delete(ctx, tx, id); err != nil {
// 		return err
// 	}

// 	if err := tx.Commit(); err != nil {
// 		return err
// 	}
// 	return nil
// }
