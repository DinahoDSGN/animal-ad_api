package models

import (
	"gorm.io/gorm"
)

type Specify struct {
	Id         uint   `json:"spec_id"`
	Name       string `json:"name"`
	BreedId    uint   `json:"breed_id"`
	Breed      *Breed `gorm:"foreignKey:BreedId"`
	Color      string `json:"color"`
	Gender     bool   `json:"gender" gorm:"type:bool"`
	Vaccinated bool   `json:"vaccinated" gorm:"type:bool"`
	Spayed     bool   `json:"spayed" gorm:"type:bool"`
	Passport   bool   `json:"passport" gorm:"type:bool"`
	Price      int16  `json:"price"`
	Profit     int16  `json:"profit"`
}

func (s *Specify) BeforeCreate(tx *gorm.DB) (err error) {
	breedGlobalPrice := s

	tx.Preload("Breed").Raw("SELECT * FROM specifies WHERE id = ?", s.Id).Find(&breedGlobalPrice)

	profit := s.Price - breedGlobalPrice.Breed.GlobalPrice
	if profit > 0 {
		s.Profit = -profit
		return
	}

	s.Profit = profit

	return
}

func (s *Specify) AfterUpdate(database *gorm.DB) error {
	database.Preload("Breed").Raw("SELECT * FROM specifies WHERE id = ?", s.Id).Find(&s)
	if s.Breed != nil {
		profit := s.Price - s.Breed.GlobalPrice
		if profit > 0 {
			database.Table("specifies").Where("id = ?", s.Id).Update("profit", profit)
			//database.Save(&s)

			return nil
		}

		database.Table("specifies").Where("id = ?", s.Id).Update("profit", profit)
		//database.Save(&s)
		return nil
	}

	return nil
}
