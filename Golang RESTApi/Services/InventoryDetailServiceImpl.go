package services

import (
	helper "RESTApi/Helper"
	exception "RESTApi/Helper/Exception"
	repository "RESTApi/Models/Repository"
	requests "RESTApi/Models/Requests"
	responses "RESTApi/Models/Responses"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type InventoryDetailServiceImpl struct {
	RepoInventoryDetails repository.InventoryDetailRepository
	RepoInventoryProduct repository.InventoryProductRepository
	DB                   *sql.DB
	RedisClient          *redis.Client
}

func NewInventoryDetailService(repoDetails repository.InventoryDetailRepository, repoInventoryProduct repository.InventoryProductRepository, db *sql.DB, redis *redis.Client) InventoryDetailService {
	return &InventoryDetailServiceImpl{
		RepoInventoryDetails: repoDetails,
		RepoInventoryProduct: repoInventoryProduct,
		DB:                   db,
		RedisClient:          redis,
	}
}

// **********************PERCOBAAN MAKE LOCKING**************************************************

// func (s *InventoryDetailServiceImpl) ChangeStock(ctx context.Context, request requests.StockChangeRequest) error {
// 	tx, err := s.DB.BeginTx(ctx, nil)
// 	if err != nil {
// 		return exception.ServiceErr(err, "error beginning transaction", "database_error")
// 	}
// 	defer helper.TxHandler(tx, err)

// 	lockId := request.ProductId // Use productId as lock identifier
// 	_, err = s.Repo.AcquireAdvisoryLock(ctx, tx, lockId)
// 	if err != nil {
// 		return exception.ServiceErr(err, "error acquiring advisory lock")
// 	}

// 	inventoryId, err := s.Repo.FindInventoryByProductId(ctx, tx, request.ProductId)
// 	if err != nil {
// 		return exception.ServiceErr(err, "error fetching inventory by product_id")
// 	}

// 	details, err := s.Repo.FindByInventoryId(ctx, tx, inventoryId)
// 	if err != nil {
// 		return exception.ServiceErr(err, "error fetching inventory details by inventory_id")
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
// 		return exception.ServiceErr(err, "error updating stock and status")
// 	}

// 	return nil
// }

func (s *InventoryDetailServiceImpl) ChangeStock(ctx context.Context, request requests.StockChangeRequest) error {

	// locking database
	lockKey := fmt.Sprintf("lock:product:%d", request.ProductId)
	lock, err := s.RedisClient.SetNX(lockKey, "locked", 10*time.Second).Result()
	if err != nil {
		return exception.ServiceErr(err, "error acquiring lock", "database_error")
	}

	if !lock {
		return exception.ServiceErr(fmt.Errorf("unable to acquire lock"), "operation already in progress", "database_error")
	}

	defer s.RedisClient.Del(lockKey)

	// Begin Transaction
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return exception.ServiceErr(err, "error beginning transaction", "database_error")
	}
	defer helper.TxHandler(tx, err)

	// Find Inventory By Product ID
	inventoryId, err := s.RepoInventoryProduct.FindInventoryByProductId(ctx, tx, request.ProductId)
	if err != nil {
		return exception.ServiceErr(err, "product not found", "not_found")
	}
	time.Sleep(100 * time.Millisecond)

	// Find Inventory Details By Inventory ID
	details, err := s.RepoInventoryDetails.FindByInventoryId(ctx, tx, inventoryId)
	if err != nil {
		return exception.ServiceErr(err, "inventory product not found", "not_found")
	}

	// Update Stock
	details.Stock = details.Stock + request.Change
	if details.Stock < 0 {
		return exception.ServiceErr(fmt.Errorf("insufficient stock: available stock is %d", details.Stock), "insufficient stock", "validation_error")
	}

	// Update Status
	switch {
	case details.Stock == 0:
		details.Status = "LOST"
	case details.Stock < 100:
		details.Status = "BAD"
	default:
		details.Status = "AVAILABLE"
	}

	// Update Stock in Database
	_, err = s.RepoInventoryDetails.UpdateStock(ctx, tx, details)
	if err != nil {
		return exception.ServiceErr(err, "failed updating stock product", "database_error")
	}

	return nil
}

func (s *InventoryDetailServiceImpl) FindInventoryDetailById(ctx context.Context, id int) (responses.InventoryDetailResponse, error) {
	tx, err := s.DB.BeginTx(ctx, helper.BeginTxHandlerExec())
	if err != nil {
		return responses.InventoryDetailResponse{}, exception.ServiceErr(err, "error beginning transaction", "database_error")
	}
	defer helper.TxHandler(tx, err)

	detail, err := s.RepoInventoryDetails.FindByInventoryId(ctx, tx, id)
	if err != nil {
		return responses.InventoryDetailResponse{}, exception.ServiceErr(err, "inventory product detail not found", "not_found")
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
