package repository

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"petcard/pkg/models"
	"strings"
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
		"INNER JOIN breeds ON animals.breed_id = breeds.id "

	if len(values) != 0 {
		queryTemplate += "WHERE"
	}

	if values["specify"] != nil {
		queryTemplate += fmt.Sprintf(" animals.type='%v' AND ", values["specify"])
	}

	if values["breed"] != nil {
		queryTemplate += fmt.Sprintf(" breeds.name='%v' AND ", values["breed"])
	}

	if values["color"] != nil {
		queryTemplate += fmt.Sprintf(" animals.color='%v' AND ", values["color"])
	}

	if values["sex"] != nil {
		queryTemplate += fmt.Sprintf(" animals.gender=%v AND ", values["sex"])
	}

	if values["vaccinated"] != nil {
		queryTemplate += fmt.Sprintf(" animals.vaccinated=%v AND ", values["vaccinated"])
	}

	if values["spayed"] != nil {
		queryTemplate += fmt.Sprintf(" animals.spayed=%v AND ", values["spayed"])
	}

	if values["passport"] != nil {
		queryTemplate += fmt.Sprintf(" animals.passport=%v AND ", values["passport"])
	}

	if values["price"] != nil {
		// input 100-800; output [100 800]
		price := strings.Split(fmt.Sprintf("%v", values["price"]), "-")

		queryTemplate += fmt.Sprintf(" animals.price >= %v AND animals.price <= %v AND ", price[0], price[1])
	}

	// trims last three words like "AND " with space at query
	queryTemplate = strings.TrimSuffix(queryTemplate, "AND ")

	fmt.Println(queryTemplate)

	query := fmt.Sprintf(queryTemplate)

	database.db.Preload("Author").Preload("Animal.Breed").
		Raw(query).
		Find(&records)

	log.Println(values)

	return records, nil
}
