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
	paramUsername := c.Params("username")
	var data models.User
	database.DB.Raw("SELECT * FROM users WHERE username = ?", paramUsername).Find(&data)

	if data.Id == 0 {
		return c.JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
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

func UserUpdateByName(c *fiber.Ctx) error {
	var data map[string]string
	paramUsername := c.Params("username")

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		Name: data["name"],
	}

	database.DB.Model(&models.User{}).Where(
		"username = ?", paramUsername,
	).
		Updates(map[string]interface{}{
			"name": user.Name,
		})

	return c.JSON(fiber.Map{
		"message": "name successfully updated",
	})
}

func UserUpdateByLastname(c *fiber.Ctx) error {
	var data map[string]string
	paramUsername := c.Params("username")

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		Lastname: data["lastname"],
	}

	database.DB.Model(&models.User{}).Where(
		"username = ?", paramUsername,
	).
		Updates(map[string]interface{}{
			"lastname": user.Lastname,
		})

	return c.JSON(fiber.Map{
		"message": "lastname successfully updated",
	})
}

func UserUpdateByUsername(c *fiber.Ctx) error {
	var data map[string]string
	paramUsername := c.Params("username")

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		Username: data["username"],
	}

	database.DB.Model(&models.User{}).Where(
		"username = ?", paramUsername,
	).
		Updates(map[string]interface{}{
			"username": user.Username,
		})

	return c.JSON(fiber.Map{
		"message": "username successfully updated",
	})
}

func UserUpdateByEmail(c *fiber.Ctx) error {
	var data map[string]string
	paramUsername := c.Params("username")

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		Email: data["email"],
	}

	database.DB.Model(&models.User{}).Where(
		"username = ?", paramUsername,
	).
		Updates(map[string]interface{}{
			"email": user.Email,
		})

	return c.JSON(fiber.Map{
		"message": "email successfully updated",
	})
}

func UserUpdateByAd(c *fiber.Ctx) error {
	var data map[string]string
	paramUsername := c.Params("username")

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	adId, _ := strconv.Atoi(data["ad_id"])

	user := models.User{
		AdId: uint(adId),
	}

	database.DB.Model(user).Where(
		"username = ?", paramUsername,
	).
		Updates(map[string]interface{}{
			"ad_id": user.AdId,
		}).Find(&user)

	if user.Id == 0 {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "cannot find ad",
		})
	}

	return c.JSON(fiber.Map{
		"message": "ad has successfully updated",
	})
}
