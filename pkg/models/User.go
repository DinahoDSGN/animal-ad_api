package models

type User struct {
	Id       uint   `json:"id" `
	Name     string `json:"name" gorm:"<-"`
	Lastname string `json:"lastname" gorm:"<-"`
	Username string `json:"username" gorm:"unique" gorm:"<-"`
	Email    string `json:"email" gorm:"unique" gorm:"<-"`
	Password string `json:"password"`
	Ad       []Ad   `gorm:"foreignKey:UserId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
