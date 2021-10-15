package models

type Breed struct {
	Id          uint   `json:"breed_id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Wool        string `json:"wool"`
	GlobalPrice int16  `json:"global_price"`
}
