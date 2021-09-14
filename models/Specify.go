package models

type Specify struct {
	Id         uint   `json:"spec_id"`
	Name       string `json:"name"`
	Breed      string `json:"breed"`
	Color      string `json:"color"`
	Gender     bool   `json:"gender" gorm:"type:bool"`
	Vaccinated bool   `json:"vaccinated" gorm:"type:bool"`
	Spayed     bool   `json:"spayed" gorm:"type:bool"`
	Passport   bool   `json:"passport" gorm:"type:bool"`
}
