package services

import (
	helper "RESTApi/Helper"
	repository "RESTApi/Models/Repository"
	requests "RESTApi/Models/Requests"
	responses "RESTApi/Models/Responses"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type InventoryDetailServiceImpl struct {
	RepoInventoryDetails repository.InventoryDetailRepository
	RepoInventoryProduct repository.InventoryProductRepository
	DB                   *sql.DB
}

func NewInventoryDetailService(repoDetails repository.InventoryDetailRepository, repoInventoryProduct repository.InventoryProductRepository, db *sql.DB) InventoryDetailService {
	return &InventoryDetailServiceImpl{
		RepoInventoryDetails: repoDetails,
		RepoInventoryProduct: repoInventoryProduct,
		DB:                   db,
	}
}

// **********************PERCOBAAN MAKE LOCKING**************************************************

// func (s *InventoryDetailServiceImpl) ChangeStock(ctx context.Context, request requests.StockChangeRequest) error {
// 	tx, err := s.DB.BeginTx(ctx, nil)
// 	if err != nil {
// 		return helper.ServiceErr(err, "error starting transaction")
// 	}
// 	defer helper.TxHandler(tx, err)

// 	lockId := request.ProductId // Use productId as lock identifier
// 	_, err = s.Repo.AcquireAdvisoryLock(ctx, tx, lockId)
// 	if err != nil {
// 		return helper.ServiceErr(err, "error acquiring advisory lock")
// 	}

// 	inventoryId, err := s.Repo.FindInventoryByProductId(ctx, tx, request.ProductId)
// 	if err != nil {
// 		return helper.ServiceErr(err, "error fetching inventory by product_id")
// 	}

// 	details, err := s.Repo.FindByInventoryId(ctx, tx, inventoryId)
// 	if err != nil {
// 		return helper.ServiceErr(err, "error fetching inventory details by inventory_id")
// 	}

// 	details.Stock += request.Change
// 	if details.Stock < 0 {
// 		return fmt.Errorf("insufficient stock: available stock is %d", details.Stock)
// 	}

// 	switch {
// 	case details.Stock == 0:
// 		details.Status = "LOST"
// 	case details.Stock < 100:
// 		details.Status = "BAD"
// 	default:
// 		details.Status = "AVAILABLE"
// 	}

// 	_, err = s.Repo.UpdateStock(ctx, tx, details)
// 	if err != nil {
// 		return helper.ServiceErr(err, "error updating stock and status")
// 	}

// 	return nil
// }

func (s *InventoryDetailServiceImpl) ChangeStock(ctx context.Context, request requests.StockChangeRequest) error {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return helper.ServiceErr(err, "error starting transaction")
	}
	defer helper.TxHandler(tx, err)

	inventoryId, err := s.RepoInventoryProduct.FindInventoryByProductId(ctx, tx, request.ProductId)
	if err != nil {
		return helper.ServiceErr(err, "error fetching inventory by product_ids")
	}
	time.Sleep(100 * time.Millisecond)

	details, err := s.RepoInventoryDetails.FindByInventoryId(ctx, tx, inventoryId)
	if err != nil {
		return helper.ServiceErr(err, "error fetching inventory details by inventory_id")
	}

	details.Stock = details.Stock + request.Change
	if details.Stock < 0 {
		return fmt.Errorf("insufficient stock: available stock is %d", details.Stock)
	}

	switch {
	case details.Stock == 0:
		details.Status = "LOST"
	case details.Stock < 100:
		details.Status = "BAD"
	default:
		details.Status = "AVAILABLE"
	}

	_, err = s.RepoInventoryDetails.UpdateStock(ctx, tx, details)
	if err != nil {
		return helper.ServiceErr(err, "error updating stock and status")
	}

	return nil
}

func (s *InventoryDetailServiceImpl) FindInventoryDetailById(ctx context.Context, id int) (responses.InventoryDetailResponse, error) {
	tx, err := s.DB.BeginTx(ctx, helper.BeginTxHandlerExec())
	if err != nil {
		return responses.InventoryDetailResponse{}, helper.ServiceErr(err, "error begin transaction")
	}
	defer helper.TxHandler(tx, err)

	detail, err := s.RepoInventoryDetails.FindByInventoryId(ctx, tx, id)
	if err != nil {
		return responses.InventoryDetailResponse{}, err
	}

	return helper.HandleInventoryDetail(detail), nil
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
