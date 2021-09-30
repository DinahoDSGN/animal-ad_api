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
		Username: data.Username,
		Password: data.Password,
	}

	database.db.Create(&user)

	return user, nil
}

func (database *AuthorizationRepo) GetUser(username string, password string) (models.User, error) {
	var data models.User

	err := database.db.Raw("SELECT * FROM users WHERE username = ? AND password = ?", username, password).Find(&data)

	if data.Id == 0 {
		return data, err.Error
	}

	return data, nil
}
