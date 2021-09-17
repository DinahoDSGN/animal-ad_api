package repository

import (
	"gorm.io/gorm"
	"petcard/internal/models"
)

type SpecRepo struct {
	db *gorm.DB
}

func NewSpecRepo(db *gorm.DB) *SpecRepo {
	return &SpecRepo{db: db}
}

func (database *SpecRepo) Create(data models.Specify) (int, error) {
	spec := models.Specify{
		Name:       data.Name,
		Breed:      data.Breed,
		Color:      data.Color,
		Gender:     data.Gender,
		Vaccinated: data.Vaccinated,
		Spayed:     data.Spayed,
		Passport:   data.Passport,
	}

	database.db.Create(&spec)

	return int(spec.Id), nil
}

func (database *SpecRepo) GetAll() ([]models.Specify, error) {
	var records []models.Specify
	database.db.Find(&records)

	return records, nil
}

func (database *SpecRepo) GetList(id int) (models.Specify, error) {
	var data models.Specify

	database.db.Raw("SELECT * FROM specifies WHERE id = ?", id).Find(&data)

	return data, nil
}

func (database *SpecRepo) Delete(id int) (models.Specify, error) {
	var data models.Specify

	err := database.db.Raw("SELECT * FROM specifies WHERE id = ?", id).Find(&data)

	if data.Id == 0 {
		return data, err.Error
	}

	database.db.Delete(&data)

	return data, nil
}

func (database *SpecRepo) Update(id int, data models.Specify) (models.Specify, error) {
	spec := models.Specify{
		Name:       data.Name,
		Breed:      data.Breed,
		Color:      data.Color,
		Gender:     data.Gender,
		Vaccinated: data.Vaccinated,
		Spayed:     data.Spayed,
		Passport:   data.Passport,
	}

	database.db.Model(&spec).Where("id = ?", id).Updates(&spec).Find(&spec)

	if spec.Id == 0 {
		return spec, nil
	}

	database.db.Save(&spec)

	return spec, nil
}
