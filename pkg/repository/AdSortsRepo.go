package repository

import (
	"fmt"
	"gorm.io/gorm"
	"petcard/pkg/models"
)

type AdSortsRepo struct {
	db *gorm.DB
}

func NewAdSortsRepo(db *gorm.DB) *AdSortsRepo {
	return &AdSortsRepo{db: db}
}

func (database AdSortsRepo) SortBy(values map[string]interface{}) ([]models.Ad, error) {
	var records []models.Ad

	queryTemplate := "SELECT * FROM ads " +
		"INNER JOIN animals ON ads.animal_id = animals.id " +
		"INNER JOIN breeds ON animals.breed_id = breeds.id " +
		"WHERE"

	if values["specify"] != "" {
		queryTemplate += fmt.Sprintf(" animals.type='%v' AND", values["specify"])
	}

	if values["breed"] != "" {
		queryTemplate += fmt.Sprintf(" breeds.name='%v' AND", values["breed"])
	}

	if values["sex"] != "" {
		queryTemplate += fmt.Sprintf(" animals.gender=%v ", values["sex"])
	}

	query := fmt.Sprintf(queryTemplate)

	database.db.Preload("Author").Preload("Animal.Breed").
		Raw(query).
		Find(&records)

	return records, nil
}
