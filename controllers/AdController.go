package controllers

import (
	"github.com/gofiber/fiber/v2"
	"petcard/database"
	"petcard/models"
	"strconv"
)

func CreateAd(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	authorId, _ := strconv.Atoi(data["author_id"])
	specifyId, _ := strconv.Atoi(data["specify_id"])

	ad := models.Ad{
		Title:       data["title"],
		Location:    data["location"],
		Description: data["description"],
		SpecifyId:   uint(specifyId),
		UserId:      uint(authorId),
	}

	database.DB.Create(&ad)

	if ad.Id == 0 {
		return c.JSON(fiber.StatusBadRequest)
	}

	database.DB.Preload("Author").Preload("Specify").Table("ads").Find(&ad)

	if ad.Author.Id == 0 {
		return c.JSON(fiber.StatusBadRequest)
	}

	return c.JSON(ad)
}

func AdGetAll(c *fiber.Ctx) error {
	var records []models.Ad
	database.DB.Preload("Specify").Preload("Author").Find(&records)

	return c.JSON(fiber.Map{
		"status": fiber.StatusFound,
		"length": len(records),
		"ads":    records,
	})
}

func AdGetByTitle(c *fiber.Ctx) error {
	paramTitle := c.Params("title")
	var data models.Ad
	database.DB.Raw("SELECT * FROM ads WHERE title = ?", paramTitle).Find(&data)

	if data.Id == 0 {
		return c.JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
		})
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusFound,
		"ad":     data,
	})
}

func AdDeleteByTitle(c *fiber.Ctx) error {
	paramTitle := c.Params("title")
	var data models.Ad

	database.DB.Raw("SELECT * FROM ads WHERE title = ?", paramTitle).Find(&data)

	if data.Id == 0 {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "cannot find ad",
		})
	}

	database.DB.Delete(&data)

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "ad successfully deleted",
	})
}
