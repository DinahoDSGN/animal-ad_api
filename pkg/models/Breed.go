package models

type Breed struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name"`
	GlobalPrice float64 `json:"global_price"`
}
