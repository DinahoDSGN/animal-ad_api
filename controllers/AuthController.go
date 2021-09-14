package controllers

import (
	"github.com/gofiber/fiber/v2"
	"petcard/database"
	"petcard/models"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil{
		return err
	}

	user := models.User{
		Name:     data["name"],
		Lastname: data["lastname"],
		Username: data["username"],
		Email:    data["email"],
		Password: data["password"],
	}

	database.DB.Create(&user)

	if user.Id == 0 {
		return c.JSON(fiber.StatusBadRequest)
	}

	return c.JSON(user)
}