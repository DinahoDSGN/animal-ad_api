package repository

import (
	"gorm.io/gorm"
	"petcard/pkg/models"
)

type BreedRepo struct {
	db *gorm.DB
}

func NewBreedRepo(db *gorm.DB) *BreedRepo {
	return &BreedRepo{db: db}
}

func (database *BreedRepo) Create(data models.Breed) (int, error) {
	breed := models.Breed{
		Name:        data.Name,
		GlobalPrice: data.GlobalPrice,
		Wool:        data.Wool,
	}

	database.db.Create(&breed)

	return int(breed.Id), nil
}

func (database *BreedRepo) GetAll() ([]models.Breed, error) {
	var records []models.Breed
	database.db.Find(&records)

	return records, nil
}

func (database *BreedRepo) GetList(id int) (models.Breed, error) {
	var data models.Breed
	database.db.Raw("SELECT * FROM breeds WHERE id = ?", id).Find(&data)

	return data, nil
}

func (database *BreedRepo) Delete(id int) (models.Breed, error) {
	var data models.Breed
	err := database.db.Raw("SELECT * FROM breeds WHERE id = ?", id).Find(&data)

	if data.Id == 0 {
		return data, err.Error
	}

	database.db.Delete(&data)

	return data, nil
}

func (database *BreedRepo) Update(id int, data models.Breed) (models.Breed, error) {
	breed := models.Breed{
		Name:        data.Name,
		GlobalPrice: data.GlobalPrice,
		Wool:        data.Wool,
	}

	database.db.Model(&breed).Where("id = ?", id).Updates(&breed).Find(&breed)

	if breed.Id == 0 {
		return breed, nil
	}

	database.db.Save(&breed)

	return breed, nil
}
