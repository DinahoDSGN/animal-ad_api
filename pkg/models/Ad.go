package models

type Ad struct {
	Id          uint    `json:"ad_id"`
	Title       string  `json:"title" gorm:"unique"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
	AnimalId    uint    `json:"animal_id"`
	Animal      *Animal `json:"animal" gorm:"foreignKey:AnimalId"`
	AuthorId    uint    `json:"author_id"`
	Author      *User   `json:"author" gorm:"foreignKey:AuthorId"`
}
