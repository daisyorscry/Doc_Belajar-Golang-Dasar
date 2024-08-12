package services

import (
	helper "RESTApi/Helper"
	entity "RESTApi/Models/Entity"
	repository "RESTApi/Models/Repository"
	requests "RESTApi/Models/Requests"
	responses "RESTApi/Models/Responses"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func (s *UserServiceImpl) FindById(ctx context.Context, id int) (responses.UserResponse, error) {
	// Mulai transaksi database
	txOption := sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  true,
	}

	tx, err := s.DB.BeginTx(ctx, &txOption)
	if err != nil {
		return responses.UserResponse{}, helper.ServiceErr(err, "error beginning transaction")
	}
	defer helper.TxHandler(tx, err)

	// Cari user berdasarkan ID
	user, err := s.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		return responses.UserResponse{}, helper.ServiceErr(err, "error finding user by id")
	}

	// Kembalikan respon
	return helper.HandleUserResponse(user), nil
}

func (s *UserServiceImpl) FindByUsername(ctx context.Context, username string) (responses.UserResponse, error) {
	// Mulai transaksi database
	txOption := sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  true,
	}

	tx, err := s.DB.BeginTx(ctx, &txOption)
	if err != nil {
		return responses.UserResponse{}, helper.ServiceErr(err, "error beginning transaction")
	}
	defer helper.TxHandler(tx, err)

	// Cari user berdasarkan username
	user, err := s.UserRepository.FindByUsername(ctx, tx, username)
	if err != nil {
		return responses.UserResponse{}, helper.ServiceErr(err, "error finding user by username")
	}

	// Kembalikan respon
	return helper.HandleUserResponse(user), nil
}

func (s *UserServiceImpl) Login(ctx context.Context, request requests.UserLoginRequest) (responses.UserResponse, string, error) {
	// Validasi input request
	err := s.Validate.Struct(request)
	if err != nil {
		return responses.UserResponse{}, "", helper.ServiceErr(err, "invalid login request")
	}

	// Mulai transaksi database
	txOption := sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  true,
	}

	tx, err := s.DB.BeginTx(ctx, &txOption)
	if err != nil {
		return responses.UserResponse{}, "", helper.ServiceErr(err, "error beginning transaction")
	}
	defer helper.TxHandler(tx, err)

	// Ambil data user berdasarkan username
	user, err := s.UserRepository.FindByUsername(ctx, tx, request.Username)
	if err != nil {
		return responses.UserResponse{}, "", helper.ServiceErr(err, "error during login")
	}

	// Verifikasi password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return responses.UserResponse{}, "", helper.ServiceErr(err, "invalid username or password")
	}

	// Generate JWT token
	token, err := helper.GenerateJWT(user.Id, user.Username)
	if err != nil {
		return responses.UserResponse{}, "", helper.ServiceErr(err, "error generating JWT")
	}

	// Kembalikan respon dan token
	return helper.HandleUserResponse(user), token, nil
}

func (s *UserServiceImpl) Register(ctx context.Context, request requests.UserRegistrationRequest) (responses.UserResponse, error) {
	// Validasi input request
	err := s.Validate.Struct(request)
	if err != nil {
		return responses.UserResponse{}, helper.ServiceErr(err, "error validating request")
	}

	// Hashing password sebelum menyimpan ke database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return responses.UserResponse{}, helper.ServiceErr(err, "error hashing password")
	}

	// begin transaction
	txOption := sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	}

	tx, err := s.DB.BeginTx(ctx, &txOption)
	if err != nil {
		return responses.UserResponse{}, helper.ServiceErr(err, "error beginning transaction")
	}
	defer helper.TxHandler(tx, err)

	// set request user to entity user
	user := entity.User{
		Username:  request.Username,
		Email:     request.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// send entitiy user to repository
	user, err = s.UserRepository.Register(ctx, tx, user)
	if err != nil {
		return responses.UserResponse{}, helper.ServiceErr(err, "error during registration")
	}

	// return response
	return helper.HandleUserResponse(user), nil
}

// user update service
func (s *UserServiceImpl) Update(ctx context.Context, request requests.UserUpdateRequest) (responses.UserResponse, error) {

	// Get the ID data that will be updated in the context value sent from JWT
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		return responses.UserResponse{}, fmt.Errorf("user ID not found in context")
	}

	// validasi input request
	err := s.Validate.Struct(request)
	if err != nil {
		return responses.UserResponse{}, helper.ServiceErr(err, "invalid update request")
	}

	// begin transaction
	txOption := sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}

	tx, err := s.DB.BeginTx(ctx, &txOption)
	if err != nil {
		return responses.UserResponse{}, helper.ServiceErr(err, "error beginning transaction")
	}
	defer helper.TxHandler(tx, err)

	// send id to repository find by id
	user, err := s.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		return responses.UserResponse{}, helper.ServiceErr(err, "user not found")
	}

	// set request to entity user
	if request.Username != "" {
		user.Username = request.Username
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	user.UpdatedAt = time.Now()

	// bcrycpt password
	if request.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return responses.UserResponse{}, helper.ServiceErr(err, "error hashing password")
		}
		user.Password = string(hashedPassword)
	}

	// sent entity user to repository update
	updatedUser, err := s.UserRepository.Update(ctx, tx, user)
	if err != nil {
		return responses.UserResponse{}, helper.ServiceErr(err, "error updating user")
	}

	// return response
	return helper.HandleUserResponse(updatedUser), nil
}
