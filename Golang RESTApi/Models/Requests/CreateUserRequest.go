package requests

type UserRegistrationRequest struct {
	Username string `validate:"required,min=1,max=20" json:"username"`
	Email    string `validate:"required,min=1,max=100" json:"email"`
	Password string `validate:"required,min=1,max=100" json:"password"`
}

type UserUpdateRequest struct {
	// Id       int    `validate:"required" json:"id"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserLoginRequest struct {
	Username string `validate:"required,min=1,max=20" json:"username"`
	Password string `validate:"required,min=1,max=100" json:"password"`
}
