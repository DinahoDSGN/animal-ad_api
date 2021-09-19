package handler

import (
	"github.com/gofiber/fiber/v2"
	"petcard/pkg/models"
)

func (h *Handler) SignIn(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) SignUp(c *fiber.Ctx) error {
	var JSONinput models.User

	if err := c.BodyParser(&JSONinput); err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	data, err := h.services.Authorization.SignUp(JSONinput)
	if err != nil {
		newErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   data,
	})
}
