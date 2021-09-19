package repository

import (
	"gorm.io/gorm"
	"petcard/pkg/models"
)

type AuthorizationRepo struct {
	db *gorm.DB
}

func NewAuthorizationRepo(db *gorm.DB) *AuthorizationRepo {
	return &AuthorizationRepo{db: db}
}

func (database *AuthorizationRepo) SignIn(data models.User) (models.User, error) {
	return models.User{}, nil
}

func (database *AuthorizationRepo) SignUp(data models.User) (models.User, error) {
	user := models.User{
		Name:     data.Name,
		Lastname: data.Lastname,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}

	database.db.Create(&user)

	return user, nil
}
