package models

type AdLocation struct {
	Id      uint   `json:"location_id"`
	Address string `json:"address"`
	Country string `json:"country"`
	Region  string `json:"region"`
	City    string `json:"city"`
	Ad      []Ad   `json:"ad" gorm:"foreignKey:AuthorId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" sql:"default:null"`
}
