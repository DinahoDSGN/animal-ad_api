package telegramModels

type Ad struct {
	AdID        int    `json:"ad_id"`
	Title       string `json:"title"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Animal      Animal `json:"animal"`
	Author      Author `json:"author"`
}
type Breed struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Wool        string      `json:"wool"`
	GlobalPrice int         `json:"global_price"`
	Animal      interface{} `json:"animal"`
}
type Animal struct {
	SpecID     int    `json:"spec_id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Color      string `json:"color"`
	Gender     bool   `json:"gender"`
	Vaccinated bool   `json:"vaccinated"`
	Spayed     bool   `json:"spayed"`
	Passport   bool   `json:"passport"`
	Breed      Breed  `json:"breed"`
	Price      int    `json:"price"`
	Profit     int    `json:"profit"`
}
type Author struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Lastname string      `json:"lastname"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Ad       interface{} `json:"ad"`
}
