package request

type CreateUserRequest struct {
	Username string `bindig:"required"`
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

type LoginRequest struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}
