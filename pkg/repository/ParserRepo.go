package repository

import (
	"github.com/gocolly/colly"
	"gorm.io/gorm"
	"petcard/pkg/models"
	"regexp"
	"strconv"
)

type ParserRepo struct {
	db *gorm.DB
}

func (database *ParserRepo) Push() error {
	var name []string
	var wool []string
	var price []int
	name, wool, price = ParseURL()

	breed := models.Breed{}

	for i := 0; i < len(name); i++ {
		breed.Id = uint(i + 1)
		breed.Name = name[i]
		breed.Type = "Кошка"
		breed.Wool = wool[i]
		breed.GlobalPrice = int16(price[i])
		database.db.Create(&breed)
	}

	return nil
}

func NewParserRepo(db *gorm.DB) *ParserRepo {
	return &ParserRepo{db: db}
}

func ParseURL() ([]string, []string, []int) {
	var linkData []string
	var nameData []string
	var priceData []int
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
				priceString := e.DOM.Next().Text()

				re := regexp.MustCompile("[0-9]+")
				numbers := re.FindAllString(priceString, -1)

				if numbers == nil {
					numbers = append(numbers, "0", "0")
				}

				num1, _ := strconv.Atoi(numbers[0])
				num2, _ := strconv.Atoi(numbers[1])

				priceAverage := (num1 + num2) / 2

				priceData = append(priceData, priceAverage)
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
