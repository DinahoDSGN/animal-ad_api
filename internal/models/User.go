package models

type User struct {
	Id       uint   `json:"id" `
	Name     string `json:"name" gorm:"<-"`
	Lastname string `json:"lastname" gorm:"<-"`
	Username string `json:"username" gorm:"unique" gorm:"<-"`
	Email    string `json:"email" gorm:"unique" gorm:"<-"`
	Password string `json:"password"`
	AdId     uint   `json:"ad_id" gorm:"default:null" gorm:"<-"`
	Ad       []Ad   `gorm:"foreignKey:UserId"`
}
