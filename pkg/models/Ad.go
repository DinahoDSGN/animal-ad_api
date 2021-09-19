package models

type Ad struct {
	Id          uint     `json:"ad_id"`
	Title       string   `json:"title" gorm:"unique"`
	Location    string   `json:"location"`
	Description string   `json:"description"`
	SpecifyId   uint     `json:"-"`
	Specify     *Specify `gorm:"foreignKey:SpecifyId"`
	UserId      uint     `json:"-"`
	Author      *User    `gorm:"foreignKey:UserId"`
}
