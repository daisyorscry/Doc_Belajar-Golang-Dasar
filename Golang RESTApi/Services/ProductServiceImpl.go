package services

import (
	helper "RESTApi/Helper"
	entity "RESTApi/Models/Entity"
	repository "RESTApi/Models/Repository"
	requests "RESTApi/Models/Requests"
	responses "RESTApi/Models/Responses"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func (s *ProductServiceImpl) Create(ctx context.Context, request requests.CreateProductRequest) (responses.ProductRespon, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return responses.ProductRespon{}, helper.ServiceErr(err, "error validating create product request")
	}

	tx, err := s.DB.BeginTx(ctx, helper.BeginTxHandlerExec())
	if err != nil {
		return responses.ProductRespon{}, helper.ServiceErr(err, "error beginning transaction")
	}
	defer helper.TxHandler(tx, err)

	product := entity.Product{
		ProductName: request.ProductName,
		ProductDesc: request.ProductDesc,
	}

	product, err = s.ProductRepository.Save(ctx, tx, product)
	if err != nil {
		return responses.ProductRespon{}, helper.ServiceErr(err, "error saving product")
	}

	return helper.HandleProductResponse(product), nil
}

func (s *ProductServiceImpl) Update(ctx context.Context, request requests.UpdateProductRequest) (responses.ProductRespon, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return responses.ProductRespon{}, helper.ServiceErr(err, "error validating update product request")
	}

	tx, err := s.DB.BeginTx(ctx, helper.BeginTxHandlerExec())
	if err != nil {
		return responses.ProductRespon{}, helper.ServiceErr(err, "error beginning transaction")
	}
	defer helper.TxHandler(tx, err)

	product, err := s.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return responses.ProductRespon{}, helper.ServiceErr(err, "error finding product by id")
	}

	product.ProductName = request.ProductName
	product.ProductDesc = request.ProductDesc

	product, err = s.ProductRepository.Update(ctx, tx, product)
	if err != nil {
		return responses.ProductRespon{}, helper.ServiceErr(err, "error updating product")
	}

	return helper.HandleProductResponse(product), nil
}

func (s *ProductServiceImpl) Delete(ctx context.Context, request int) error {

	tx, err := s.DB.BeginTx(ctx, helper.BeginTxHandlerExec())
	if err != nil {
		return helper.ServiceErr(err, "error beginning transaction")
	}
	defer helper.TxHandler(tx, err)

	product, err := s.ProductRepository.FindById(ctx, tx, request)
	if err != nil {
		return helper.ServiceErr(err, "error finding product by id")
	}

	err = s.ProductRepository.Delete(ctx, tx, product)
	if err != nil {
		return helper.ServiceErr(err, "error deleting product")
	}

	return nil
}

func (s *ProductServiceImpl) FindById(ctx context.Context, request int) (responses.ProductRespon, error) {

	tx, err := s.DB.BeginTx(ctx, helper.BeginTxHandlerQuery())
	if err != nil {
		return responses.ProductRespon{}, helper.ServiceErr(err, "error beginning transaction")
	}
	defer helper.TxHandler(tx, err)

	product, err := s.ProductRepository.FindById(ctx, tx, request)
	if err != nil {
		return responses.ProductRespon{}, helper.ServiceErr(err, "error finding product by id")
	}

	return helper.HandleProductResponse(product), nil
}

func (s *ProductServiceImpl) FindAll(ctx context.Context) ([]responses.ProductRespon, error) {

	tx, err := s.DB.BeginTx(ctx, helper.BeginTxHandlerQuery())
	if err != nil {
		return nil, helper.ServiceErr(err, "error beginning transaction")
	}
	defer helper.TxHandler(tx, err)

	products, err := s.ProductRepository.FindAll(ctx, tx)
	if products == nil {
		return nil, helper.ServiceErr(err, "error finding all products")
	}

	return helper.HandleProductResponses(products), nil
}
