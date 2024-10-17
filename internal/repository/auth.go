package repository

import (
	"task-manager/database/models"
	"task-manager/pkg/utils"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Save(user models.User) error
	FindByEmail(email string) (models.User, error)
}

type authRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepositoryImpl {
	return &authRepositoryImpl{DB: db}
}

func (a *authRepositoryImpl) Save(user models.User) error {
	result := a.DB.Create(&user)
	return utils.ReturnError(result.Error)
}

func (a *authRepositoryImpl) FindByEmail(email string) (models.User, error) {
	var user models.User
	result := a.DB.Where("email = ?", email).First(&user)
	return user, utils.ReturnError(result.Error)
}
