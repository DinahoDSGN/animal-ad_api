package repository

import (
	"gorm.io/gorm"
	"petcard/pkg/models"
)

type AdLocationRepo struct {
	db *gorm.DB
}

func NewAdLocationRepo(db *gorm.DB) *AdLocationRepo {
	return &AdLocationRepo{db: db}
}

func (database *AdLocationRepo) Create(data models.AdLocation) (models.AdLocation, error) {
	adLocation := models.AdLocation{
		Address: data.Address,
		Country: data.Country,
		Region:  data.Region,
		City:    data.City,
	}

	database.db.Create(&adLocation)

	database.db.Preload("Author").Preload("Animal.Breed").Table("ad_locations").Find(&adLocation)

	return adLocation, nil
}

func (database *AdLocationRepo) GetAll() ([]models.AdLocation, error) {
	var records []models.AdLocation

	database.db.Preload("Ad").Find(&records)

	return records, nil
}

func (database *AdLocationRepo) GetList(id int) (models.AdLocation, error) {
	var data models.AdLocation

	database.db.Preload("Ad").Raw("SELECT * FROM ad_locations WHERE id = ?", id).Find(&data)

	return data, nil
}
