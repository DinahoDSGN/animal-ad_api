package models

type Breed struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	GlobalPrice int16  `json:"global_price"`
	Wool        string `json:"wool"`
}
