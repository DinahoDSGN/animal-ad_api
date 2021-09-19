package repository

import (
	"gorm.io/gorm"
	"petcard/pkg/models"
)

type AnimalRepo struct {
	db *gorm.DB
}

func NewAnimalRepo(db *gorm.DB) *AnimalRepo {
	return &AnimalRepo{db: db}
}

func (database *AnimalRepo) Create(data models.Animal) (int, error) {
	spec := models.Animal{
		Name:       data.Name,
		BreedId:    data.BreedId,
		Color:      data.Color,
		Gender:     data.Gender,
		Vaccinated: data.Vaccinated,
		Spayed:     data.Spayed,
		Passport:   data.Passport,
		Price:      data.Price,
	}

	database.db.Create(&spec)

	return int(spec.Id), nil
}

func (database *AnimalRepo) GetAll() ([]models.Animal, error) {
	var records []models.Animal
	database.db.Preload("Breed").Find(&records)

	return records, nil
}

func (database *AnimalRepo) GetList(id int) (models.Animal, error) {
	var data models.Animal

	database.db.Preload("Breed").Raw("SELECT * FROM animals WHERE id = ?", id).Find(&data)

	return data, nil
}

func (database *AnimalRepo) Delete(id int) (models.Animal, error) {
	var data models.Animal

	err := database.db.Raw("SELECT * FROM animals WHERE id = ?", id).Find(&data)

	if data.Id == 0 {
		return data, err.Error
	}

	database.db.Delete(&data)

	return data, nil
}

func (database *AnimalRepo) Update(id int, data models.Animal) (models.Animal, error) {
	spec := models.Animal{
		Name:       data.Name,
		BreedId:    data.BreedId,
		Color:      data.Color,
		Gender:     data.Gender,
		Vaccinated: data.Vaccinated,
		Spayed:     data.Spayed,
		Passport:   data.Passport,
		Price:      data.Price,
	}

	database.db.Model(&spec).Where("id = ?", id).Updates(&spec).Find(&spec)

	if spec.Id == 0 {
		return spec, nil
	}

	database.db.Save(&spec)

	return spec, nil
}
