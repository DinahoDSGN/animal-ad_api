package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *fiber.Ctx, status int, message string) error {
	logrus.Error(message)
	return c.JSON(fiber.Map{
		"status":  status,
		"message": message,
	})
}
