package models

type User struct {
	Id       uint   `json:"id" `
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Ad       []Ad   `json:"ad" gorm:"foreignKey:AuthorId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
