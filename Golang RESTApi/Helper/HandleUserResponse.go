package helper

import (
	entity "RESTApi/Models/Entity"
	responses "RESTApi/Models/Responses"
)

func HandleUserResponse(user entity.User) responses.UserResponse {
	return responses.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
