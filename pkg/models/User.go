package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username"`
	Email    string `json:"email" gorm:"unique" sql:"default:null"`
	Password string `json:"password"`
	Ad       []Ad   `json:"ad" gorm:"foreignKey:AuthorId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
