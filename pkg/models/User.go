package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" sql:"default:null"`
	Lastname string `json:"lastname" sql:"default:null"`
	Username string `json:"username" sql:"default:null"`
	Email    string `json:"email" sql:"default:null"`
	Password string `json:"password"`
	Ad       []Ad   `json:"ad" gorm:"foreignKey:AuthorId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" sql:"default:null"`
}
