package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Ad struct {
	Id           uint        `json:"ad_id"`
	Title        string      `json:"title" gorm:"unique"`
	Location     string      `json:"location"`
	AdLocationId uint        `json:"ad_location_id"`
	AdLocation   *AdLocation `json:"ad_location" gorm:"foreignKey:AdLocationId"`
	Description  string      `json:"description"`
	Views        uint        `json:"views"`
	Verified     bool        `json:"verified" sql:"default:true" gorm:"type:bool"`
	AnimalId     uint        `json:"animal_id"`
	Animal       *Animal     `json:"animal" gorm:"foreignKey:AnimalId"`
	AuthorId     uint        `json:"author_id"`
	Author       *User       `json:"author" gorm:"foreignKey:AuthorId"`
}

func (s *Ad) BeforeCreate(database *gorm.DB) error {
	temp := s

	database.Raw("SELECT * FROM ads WHERE id = ?", s.Id).Find(&temp)

	fmt.Println(temp.AuthorId)

	if temp.AuthorId == 0 {
		s.Verified = false
		return nil
	}

	s.Verified = true

	fmt.Println(s.Verified)

	return nil
}
