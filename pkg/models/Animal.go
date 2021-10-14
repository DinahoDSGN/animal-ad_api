package models

import (
	"gorm.io/gorm"
)

type Animal struct { // Animal struct
	Id         uint   `json:"animal_id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Color      string `json:"color"`
	Gender     bool   `json:"gender" gorm:"type:bool"`
	Vaccinated bool   `json:"vaccinated" gorm:"type:bool"`
	Spayed     bool   `json:"spayed" gorm:"type:bool"`
	Passport   bool   `json:"passport" gorm:"type:bool"`
	BreedId    uint   `json:"breed_id"`
	Breed      *Breed `json:"breed" gorm:"foreignKey:BreedId"`
	Price      int16  `json:"price"`
	Profit     int16  `json:"profit"`
}

func (s *Animal) BeforeCreate(tx *gorm.DB) (err error) {
	animal := s

	tx.Preload("Breed").Raw("SELECT * FROM animals WHERE id = ?", s.Id).Find(&animal)
	if s.Breed != nil {
		s.Type = animal.Breed.Type

		profit := s.Price - animal.Breed.GlobalPrice
		if profit > 0 {
			s.Profit = -profit
			return
		}

		s.Profit = profit
	}

	return
}

func (s *Animal) AfterUpdate(database *gorm.DB) error {
	database.Preload("Breed").Raw("SELECT * FROM animals WHERE id = ?", s.Id).Find(&s)
	if s.Breed != nil {
		profit := s.Price - s.Breed.GlobalPrice
		if profit > 0 {
			database.Table("animals").Where("id = ?", s.Id).Update("profit", profit)
			return nil
		}

		database.Table("animals").Where("id = ?", s.Id).Update("profit", profit)

		return nil
	}

	return nil
}
