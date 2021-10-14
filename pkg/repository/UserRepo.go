package repository

import (
	"gorm.io/gorm"
	"petcard/pkg/models"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (database *UserRepo) GetAll() ([]models.User, error) {
	var records []models.User
	database.db.Preload("Ad").Find(&records)

	return records, nil
}

func (database *UserRepo) GetList(id int) (models.User, error) {
	var data models.User

	database.db.Preload("Ad").Raw("SELECT * FROM users WHERE id = ?", id).Find(&data)

	return data, nil
}

func (database *UserRepo) Delete(id int) (models.User, error) {
	var data models.User

	err := database.db.Raw("SELECT * FROM users WHERE id = ?", id).Find(&data)

	if data.Id == 0 {
		return data, err.Error
	}

	database.db.Delete(&data)

	return data, nil
}

func (database *UserRepo) UpdateRating(id int, data models.User) (float32, error) {
	user := models.User{
		Rating: data.Rating,
	}

	database.db.Model(&user).Where("id = ?", id).Updates(&user).Find(&user)

	if user.Id == 0 {
		return 0, nil
	}

	database.db.Save(&user)

	return user.Rating, nil
}

func (database *UserRepo) Update(id int, data models.User) (models.User, error) {
	user := models.User{
		Name:     data.Name,
		Lastname: data.Lastname,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Ad:       data.Ad,
	}

	database.db.Model(&user).Where("id = ?", id).Updates(&user).Find(&user)

	if user.Id == 0 {
		return user, nil
	}

	database.db.Save(&user)

	return user, nil
}
