package controllers

import (
	"github.com/gofiber/fiber/v2"
	"petcard/database"
	"petcard/models"
	"strconv"
)

func UserGetAll(c *fiber.Ctx) error {
	var records []models.User
	database.DB.Preload("Ad").Find(&records)

	return c.JSON(fiber.Map{
		"status": fiber.StatusFound,
		"users":  records,
	})
}

func UserGetByUsername(c *fiber.Ctx) error {
	var data models.User
	paramUsername := c.Params("username")
	database.DB.Raw("SELECT * FROM users WHERE username = ?", paramUsername).Find(&data)

	if data.Id == 0 {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "user does not exist",
		})
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusFound,
		"user":   data,
	})
}

func UserDeleteByUsername(c *fiber.Ctx) error {
	paramUsername := c.Params("username")
	var data models.User

	database.DB.Where("username = ?", paramUsername).Find(&data)

	if data.Id == 0 {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "cannot find user",
		})
	}

	database.DB.Delete(&data)

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "user successfully deleted",
	})
}

func UserUpdate(c *fiber.Ctx) error {
	var data map[string]string
	paramUsername := c.Params("username")

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	adId, _ := strconv.Atoi(data["ad_id"])

	user := models.User{
		Name:     data["name"],
		Lastname: data["lastname"],
		Username: data["username"],
		Email:    data["email"],
		AdId:     uint(adId),
	}

	database.DB.Model(&user).Where("username = ?", paramUsername).Updates(&user).Find(&user)

	if user.Id == 0 {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "cannot find user by username",
		})
	}

	database.DB.Save(&user)

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "user successfully updated",
		"data":    user,
	})

}
