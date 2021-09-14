package models

type Ad struct {
	Id          uint     `json:"ad_id"`
	Title       string   `json:"title"`
	Location    string   `json:"location"`
	Description string   `json:"description"`
	SpecifyId   uint     `json:"specify_id"`
	Specify     *Specify `gorm:"foreignKey:SpecifyId"`
	UserId      uint     `json:"author_id"`
	Author      *User    `gorm:"foreignKey:UserId"`
}
