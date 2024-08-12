package services

import (
	helper "RESTApi/Helper"
	entity "RESTApi/Models/Entity"
	repository "RESTApi/Models/Repository"
	requests "RESTApi/Models/Requests"
	responses "RESTApi/Models/Responses"
	"context"
	"database/sql"
)

type InventoryProductServiceImpl struct {
	Repo repository.InventoryProductRepository
	DB   *sql.DB
}

func NewInventoryProductService(repo repository.InventoryProductRepository, db *sql.DB) InventoryProductService {
	return &InventoryProductServiceImpl{
		Repo: repo,
		DB:   db,
	}
}

func (s *InventoryProductServiceImpl) Create(ctx context.Context, request requests.CreateInventoryProductRequest) (responses.InventoryProductResponse, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return responses.InventoryProductResponse{}, err
	}
	defer helper.TxHandler(tx, err)

	product := entity.InventoryProduct{
		ProductId: request.ProductId,
		Price:     request.Price,
	}

	newProduct, err := s.Repo.Create(ctx, tx, product)
	if err != nil {
		return responses.InventoryProductResponse{}, err
	}

	return helper.HandleProductInventory(newProduct), nil
}

func (s *InventoryProductServiceImpl) FindById(ctx context.Context, id int) (responses.InventoryProductResponse, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return responses.InventoryProductResponse{}, err
	}
	defer helper.TxHandler(tx, err)

	product, err := s.Repo.FindById(ctx, tx, id)
	if err != nil {
		return responses.InventoryProductResponse{}, err
	}

	return helper.HandleProductInventory(product), nil
}

func (s *InventoryProductServiceImpl) FindAll(ctx context.Context) ([]responses.InventoryProductResponse, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	products, err := s.Repo.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	return helper.HandleProductInventories(products), nil
}

func (s *InventoryProductServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.TxHandler(tx, err)

	if err := s.Repo.Delete(ctx, tx, id); err != nil {
		return err
	}

	return nil
}
