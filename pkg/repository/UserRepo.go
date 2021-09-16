package repository

import (
	"gorm.io/gorm"
	"petcard/internal/models"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (database *UserRepo) GetAll() ([]models.User, error) {
	var records []models.User
	database.db.Find(&records)

	return records, nil
}

func (database *UserRepo) GetList(id int) (models.User, error) {
	var data models.User

	database.db.Raw("SELECT * FROM users WHERE id = ?", id).Find(&data)

	return data, nil
}

func (database *UserRepo) Delete(id int) error {
	var data models.User

	database.db.Raw("SELECT * FROM users WHERE id = ?", id).Find(&data)

	database.db.Delete(&data)

	return nil
}

func (database *UserRepo) Update(id int, data models.User) error {
	user := models.User{
		Name:     data.Name,
		Lastname: data.Lastname,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		AdId:     data.AdId,
	}

	database.db.Model(&user).Where("id = ?", id).Updates(&user).Find(&user)

	database.db.Save(&user)

	return nil
}
