package models

type User struct {
	Id       uint    `json:"user_id"`
	Name     string  `json:"name" sql:"default:null"`
	Lastname string  `json:"lastname" sql:"default:null"`
	Username string  `json:"username" sql:"default:null"`
	Email    string  `json:"email" sql:"default:null"`
	Password string  `json:"-"`
	Ad       []Ad    `json:"ad" gorm:"foreignKey:AuthorId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" sql:"default:null"`
	Rating   float32 `json:"rating" sql:"default:0.1"`
}
