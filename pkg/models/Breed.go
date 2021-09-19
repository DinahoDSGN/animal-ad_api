package models

type Breed struct {
	Id          uint     `json:"id"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Wool        string   `json:"wool"`
	GlobalPrice int16    `json:"global_price"`
	Animal      []Animal `gorm:"foreignKey:BreedId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
