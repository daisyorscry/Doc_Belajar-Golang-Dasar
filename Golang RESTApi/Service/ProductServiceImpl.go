package service

import (
	helper "RESTApi/Helper"
	entity "RESTApi/Models/Entity"
	repository "RESTApi/Models/Repository"
	requests "RESTApi/Models/Requests"
	responses "RESTApi/Models/Responses"
	"context"
	"database/sql"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
}

// func
func (s *ProductServiceImpl) Create(ctx context.Context, request requests.ProductRequest) responses.ProductRespon {
	txOption := sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	}

	tx, err := s.DB.BeginTx(ctx, &txOption)
	helper.DatabaseErr(err)

	defer helper.TxHandler(tx, err)

	product := entity.Product{
		ProductName: request.ProductName,
		ProductDesc: request.ProductDesc,
	}

	product, err = s.ProductRepository.Save(ctx, tx, product)
	helper.DatabaseErr(err)

	return helper.HandleProductResponse(product)
}

func (s *ProductServiceImpl) Update(ctx context.Context, request requests.ProductRequest) responses.ProductRespon {
	txOption := sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	}

	tx, err := s.DB.BeginTx(ctx, &txOption)
	helper.DatabaseErr(err)
	defer helper.TxHandler(tx, err)

	product, err := s.ProductRepository.FindById(ctx, tx, int(request.Id))
	helper.DatabaseErr(err)

	product.ProductName = request.ProductName
	product.ProductDesc = request.ProductDesc

	product, err = s.ProductRepository.Update(ctx, tx, product)
	helper.DatabaseErr(err)

	return helper.HandleProductResponse(product)
}

func (s ProductServiceImpl) Delete(ctx context.Context, request requests.ProductRequest) {
	txOption := sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	}

	tx, err := s.DB.BeginTx(ctx, &txOption)
	helper.DatabaseErr(err)
	defer helper.TxHandler(tx, err)

	product, err := s.ProductRepository.FindById(ctx, tx, int(request.Id))
	helper.DatabaseErr(err)

	s.ProductRepository.Delete(ctx, tx, product)
	helper.DatabaseErr(err)
}
func (s ProductServiceImpl) FindById(ctx context.Context, request requests.ProductRequest) responses.ProductRespon {
	txOption := sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  true,
	}

	tx, err := s.DB.BeginTx(ctx, &txOption)
	helper.DatabaseErr(err)
	defer helper.TxHandler(tx, err)

	product, err := s.ProductRepository.FindById(ctx, tx, int(request.Id))
	helper.DatabaseErr(err)

	return helper.HandleProductResponse(product)

}

func (s ProductServiceImpl) FindAll(ctx context.Context) []responses.ProductRespon {
	txOption := sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  true,
	}

	tx, err := s.DB.BeginTx(ctx, &txOption)
	helper.DatabaseErr(err)
	defer helper.TxHandler(tx, err)

	products := s.ProductRepository.FindAll(ctx, tx)

	return helper.HandleProductResponses(products)

}
