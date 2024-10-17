package service

import (
	"task-manager/database/models"
	"task-manager/internal/data/request"
	"task-manager/internal/repository"
	"task-manager/pkg/utils"

	"github.com/go-playground/validator/v10"
)

type AuthService interface {
	Create(user request.CreateUserRequest) error
	FindByEmail(email string) (models.User, error)
}

type AuthSeriveImpl struct {
	AuthRepository repository.AuthRepository
	Validate       *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, validate *validator.Validate) AuthService {
	return &AuthSeriveImpl{
		AuthRepository: authRepository,
		Validate:       validate,
	}
}

func (a *AuthSeriveImpl) Create(user request.CreateUserRequest) error {
	err := a.Validate.Struct(user)
	if err != nil {
		utils.ReturnError(err)
	}

	var newUser = models.User{
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.Password,
	}

	return utils.ReturnError(a.AuthRepository.Save(newUser))
}

func (a *AuthSeriveImpl) FindByEmail(email string) (models.User, error) {
	return a.AuthRepository.FindByEmail(email)
}
