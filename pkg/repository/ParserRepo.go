package repository

import (
	"github.com/gocolly/colly"
	"gorm.io/gorm"
	"petcard/pkg/models"
)

type ParserRepo struct {
	db *gorm.DB
}

func (database *ParserRepo) Push() error {
	var name []string
	var wool []string
	var price []string
	name, wool, price = ParseURL()

	breed := models.Breed{}

	for i := 0; i < len(name); i++ {
		breed.Id = uint(i + 1)
		breed.Name = name[i]
		breed.Wool = wool[i]
		breed.GlobalPrice = price[i]
		database.db.Create(&breed)
	}

	return nil
}

func NewParserRepo(db *gorm.DB) *ParserRepo {
	return &ParserRepo{db: db}
}

func ParseURL() ([]string, []string, []string) {
	var linkData []string
	var nameData []string
	var priceData []string
	var woolData []string

	mainPage := colly.NewCollector()

	mainPage.OnHTML(".page-dog-breeds__list-item-title", func(e *colly.HTMLElement) {
		name := e.Text
		nameData = append(nameData, name)
		link, _ := e.DOM.Attr("href")
		linkData = append(linkData, link)
	})

	mainPage.Visit("https://petsi.net/cat-breeds/")

	for i := 0; i < len(linkData); i++ {
		infoPage := colly.NewCollector()
		infoPage.OnHTML(".breed-view__table-info tr td", func(e *colly.HTMLElement) {
			text := e.Text
			if text == "Цена" {
				price := e.DOM.Next().Text()
				priceData = append(priceData, price)
			}
			if text == "Тип шерсти" {
				wool := e.DOM.Next().Text()
				woolData = append(woolData, wool)
			}
		})
		infoPage.Visit("https://petsi.net" + linkData[i])
	}

	return nameData, woolData, priceData
}
