package controllers

import (
	"github.com/gofiber/fiber/v2"
	"petcard/internal/database"
	"petcard/internal/models"
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

func SpecifyGetAll(c *fiber.Ctx) error {
	var records []models.Specify
	database.DB.Find(&records)

	return c.JSON(fiber.Map{
		"status":    fiber.StatusFound,
		"specifies": records,
	})
}

func SpecifyGetById(c *fiber.Ctx) error {
	var data models.Specify
	paramId := c.Params("id")
	database.DB.Raw("SELECT * FROM specifies WHERE id = ?", paramId).Find(&data)

	if data.Id == 0 {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "specify does not exist",
		})
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusFound,
		"user":   data,
	})
}

func SpecifyDeleteById(c *fiber.Ctx) error {
	var data models.Specify
	paramId := c.Params("id")

	database.DB.Where("id = ?", paramId).Find(&data)

	if data.Id == 0 {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "cannot find specify",
		})
	}

	database.DB.Delete(&data)

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "specify successfully deleted",
	})

}

func SpecifyUpdate(c *fiber.Ctx) error {
	var data map[string]string
	paramId := c.Params("id")

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	gender, _ := strconv.ParseBool(data["gender"])
	vaccinated, _ := strconv.ParseBool(data["vaccinated"])
	spayed, _ := strconv.ParseBool(data["spayed"])
	passport, _ := strconv.ParseBool(data["passport"])

	specify := models.Specify{
		Name:       data["name"],
		Breed:      data["breed"],
		Color:      data["color"],
		Gender:     gender,
		Vaccinated: vaccinated,
		Spayed:     spayed,
		Passport:   passport,
	}

	//database.DB.Raw("SELECT * FROM ads WHERE id = ?", paramId).Find(&specify)
	database.DB.Model(&specify).Where("id = ?", paramId).Updates(&specify).Find(&specify)

	if specify.Id == 0 {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "cannot find ad by title",
		})
	}

	database.DB.Save(&specify)

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "title successfully updated",
		"data":    specify,
	})
}
