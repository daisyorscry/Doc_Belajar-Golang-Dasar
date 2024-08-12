package services

import (
	requests "RESTApi/Models/Requests"
	responses "RESTApi/Models/Responses"
	"context"
)

type UserService interface {
	Login(ctx context.Context, request requests.UserLoginRequest) (responses.UserResponse, string, error)
	Register(ctx context.Context, request requests.UserRegistrationRequest) (responses.UserResponse, error)
	FindById(ctx context.Context, id int) (responses.UserResponse, error)
	FindByUsername(ctx context.Context, username string) (responses.UserResponse, error)
	Update(ctx context.Context, request requests.UserUpdateRequest) (responses.UserResponse, error)
}
