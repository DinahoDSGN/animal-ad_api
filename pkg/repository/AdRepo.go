package repository

import (
	"gorm.io/gorm"
	"petcard/pkg/models"
)

type AdRepo struct {
	db *gorm.DB
}

func NewAdRepo(db *gorm.DB) *AdRepo {
	return &AdRepo{db: db}
}

func (database *AdRepo) Create(data models.Ad) (models.Ad, error) {
	ad := models.Ad{
		Title:       data.Title,
		Location:    data.Location,
		Description: data.Description,
		SpecifyId:   data.SpecifyId,
		UserId:      data.UserId,
	}

	ad.Location = NewLocationData(ad.Location)

	database.db.Create(&ad)

	database.db.Preload("Author").Preload("Specify").Table("ads").Find(&ad)

	return ad, nil
}

func (database *AdRepo) GetAll() ([]models.Ad, error) {
	var records []models.Ad
	database.db.Preload("Specify").Preload("Author").Find(&records)

	return records, nil
}

func (database *AdRepo) GetList(id int) (models.Ad, error) {
	var data models.Ad
	database.db.Raw("SELECT * FROM ads WHERE id = ?", id).Find(&data)

	return data, nil
}

func (database *AdRepo) Delete(id int) error {
	var data models.Ad

	database.db.Raw("SELECT * FROM ads WHERE id = ?", id).Find(&data)

	database.db.Delete(&data)

	return nil
}

func (database *AdRepo) Update(id int, data models.Ad) error {
	ad := models.Ad{
		Title:       data.Title,
		Location:    data.Location,
		Description: data.Description,
		SpecifyId:   data.SpecifyId,
		UserId:      data.UserId,
	}

	database.db.Model(&ad).Where("id = ?", id).Updates(&ad).Find(&ad)

	database.db.Save(&ad)

	return nil
}
