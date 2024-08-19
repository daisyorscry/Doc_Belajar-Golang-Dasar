package services

import (
	helper "RESTApi/Helper"
	exception "RESTApi/Helper/Exception"
	entity "RESTApi/Models/Entity"
	repository "RESTApi/Models/Repository"
	requests "RESTApi/Models/Requests"
	responses "RESTApi/Models/Responses"
	"context"
	"database/sql"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository          repository.ProductRepository
	InventoryRepository        repository.InventoryProductRepository
	InventoryDetailsRepository repository.InventoryDetailRepository
	UserRepository             repository.UserRepository
	DB                         *sql.DB
	Validate                   *validator.Validate
}

func NewProductService(repositoryProduct repository.ProductRepository, repositoryInventory repository.InventoryProductRepository, repositoryDetails repository.InventoryDetailRepository, repositoryUser repository.UserRepository, db *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository:          repositoryProduct,
		InventoryRepository:        repositoryInventory,
		InventoryDetailsRepository: repositoryDetails,
		UserRepository:             repositoryUser,
		DB:                         db,
		Validate:                   validate,
	}
}

// only create product no inventory
func (s *ProductServiceImpl) Create(ctx context.Context, request requests.CreateProductRequest) (responses.ProductRespon, error) {
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		return responses.ProductRespon{}, exception.ServiceErr(fmt.Errorf("user ID not found"), "user ID not found in context", "unauthorized")
	}

	err := s.Validate.Struct(request)
	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "error validating create product request", "validation_error")
	}

	tx, err := s.DB.BeginTx(ctx, helper.BeginTxHandlerExec())
	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "error beginning transaction", "database_error")
	}
	defer helper.TxHandler(tx, err)

	product := entity.Product{
		ProductName: request.ProductName,
		ProductDesc: request.ProductDesc,
	}

	product, err = s.ProductRepository.Save(ctx, tx, product, userId)
	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "error creating product", "database_error")
	}

	user, err := s.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "user not found", "not_found")
	}

	return helper.HandleProductResponse(product, user), nil
}

// update product
func (s *ProductServiceImpl) Update(ctx context.Context, request requests.UpdateProductRequest) (responses.ProductRespon, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "error validating create product request", "validation_error")
	}

	tx, err := s.DB.BeginTx(ctx, helper.BeginTxHandlerExec())
	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "error beginning transaction", "database_error")
	}
	defer helper.TxHandler(tx, err)

	product, err := s.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "user not found", "not_found")
	}

	product.ProductName = request.ProductName
	product.ProductDesc = request.ProductDesc

	product, err = s.ProductRepository.Update(ctx, tx, product)
	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "update product failed", "database_error")
	}

	return helper.HandleProductResponse(product, entity.User{}), nil
}

// create product with inventory and stock
func (s *ProductServiceImpl) CreateProductWithInventoryDetails(ctx context.Context, request requests.CreateProductRequest) (responses.ProductRespon, error) {
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		return responses.ProductRespon{}, exception.ServiceErr(fmt.Errorf("user ID not found"), "user ID not found in context", "unauthorized")
	}

	err := s.Validate.Struct(request)
	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "error validating create product request", "validation_error")
	}

	tx, err := s.DB.BeginTx(ctx, helper.BeginTxHandlerExec())
	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "error beginning transaction", "database_error")
	}
	defer helper.TxHandler(tx, err)

	product := entity.Product{
		ProductName: request.ProductName,
		ProductDesc: request.ProductDesc,
	}

	product, err = s.ProductRepository.Save(ctx, tx, product, userId)
	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "failed creating product", "database_error")
	}

	inventoryProduct := entity.InventoryProduct{
		ProductId: int(product.Id),
		Price:     0.0,
	}

	inventories, err := s.InventoryRepository.Create(ctx, tx, inventoryProduct)
	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "failed creating product", "database_error")
	}

	detail := entity.InventoryDetail{
		InventoryProductId: inventories.Id,
		Stock:              0,
		Status:             "BAD",
	}

	_, err = s.InventoryDetailsRepository.Create(ctx, tx, detail)

	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "failed creating product", "database_error")
	}

	user, err := s.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "user not found", "not_found")
	}

	return helper.HandleProductResponse(product, user), nil
}

// delete product by id
func (s *ProductServiceImpl) Delete(ctx context.Context, request int) error {

	tx, err := s.DB.BeginTx(ctx, helper.BeginTxHandlerExec())
	if err != nil {
		return exception.ServiceErr(err, "error beginning transaction", "database_error")
	}
	defer helper.TxHandler(tx, err)

	product, err := s.ProductRepository.FindById(ctx, tx, request)
	if err != nil {
		return exception.ServiceErr(err, "product not found", "not_found")
	}

	err = s.ProductRepository.Delete(ctx, tx, product)
	if err != nil {
		return exception.ServiceErr(err, "failed delete product", "database_error")
	}

	return nil
}

// find product by id
func (s *ProductServiceImpl) FindById(ctx context.Context, request int) (responses.ProductRespon, error) {

	userId, ok := ctx.Value("userId").(int)
	if !ok {
		return responses.ProductRespon{}, exception.ServiceErr(fmt.Errorf("user ID not found"), "user ID not found in context", "unauthorized")
	}

	tx, err := s.DB.BeginTx(ctx, helper.BeginTxHandlerQuery())
	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "error beginning transaction", "database_error")
	}
	defer helper.TxHandler(tx, err)

	product, err := s.ProductRepository.FindById(ctx, tx, request)
	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "product not found", "not_found")
	}

	user, err := s.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		return responses.ProductRespon{}, exception.ServiceErr(err, "user not found", "not_found")
	}

	return helper.HandleProductResponse(product, user), nil
}

// find all product
func (s *ProductServiceImpl) FindAll(ctx context.Context) ([]responses.ProductRespon, error) {

	tx, err := s.DB.BeginTx(ctx, helper.BeginTxHandlerQuery())
	if err != nil {
		return nil, exception.ServiceErr(err, "error beginning transaction", "database_error")
	}
	defer helper.TxHandler(tx, err)

	products, err := s.ProductRepository.FindAll(ctx, tx)
	if products == nil {
		return nil, exception.ServiceErr(err, "failed get product", "not_found")
	}

	return helper.HandleProductResponses(products), nil
}

// find product details
func (s *ProductServiceImpl) FindProductDetail(ctx context.Context, request int) (responses.ProductDetailRespon, error) {

	tx, err := s.DB.BeginTx(ctx, helper.BeginTxHandlerExec())
	if err != nil {
		return responses.ProductDetailRespon{}, exception.ServiceErr(err, "error beginning transaction", "database_error")
	}
	defer helper.TxHandler(tx, err)

	product, err := s.ProductRepository.FindById(ctx, tx, request)
	if err != nil {
		return responses.ProductDetailRespon{}, exception.ServiceErr(err, "product not found", "not_found")
	}

	var productid int64 = product.Id

	inventory, err := s.InventoryRepository.FindInventoryByProductId(ctx, tx, int(productid))
	if err != nil {
		return responses.ProductDetailRespon{}, exception.ServiceErr(err, "inventory product not found", "not_found")
	}

	detail, err := s.InventoryDetailsRepository.FindByInventoryId(ctx, tx, inventory)
	if err != nil {
		return responses.ProductDetailRespon{}, exception.ServiceErr(err, "inventory product detail not found", "not_found")
	}

	return helper.HandleProductDetailResponse(product, detail), nil
}
