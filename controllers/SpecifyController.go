package controllers

import (
	"github.com/gofiber/fiber/v2"
	"petcard/database"
	"petcard/models"
	"strconv"
)

func CreateSpecify(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	gender, _ := strconv.ParseBool(data["gender"])
	vaccinated, _ := strconv.ParseBool(data["vaccinated"])
	spayed, _ := strconv.ParseBool(data["spayed"])
	passport, _ := strconv.ParseBool(data["passport"])

	spec := models.Specify{
		Name:       data["name"],
		Breed:      data["breed"],
		Color:      data["color"],
		Gender:     gender,
		Vaccinated: vaccinated,
		Spayed:     spayed,
		Passport:   passport,
	}

	database.DB.Create(&spec)

	//database.DB.Table("specifies").Find(&spec)

	return c.JSON(spec)

}
