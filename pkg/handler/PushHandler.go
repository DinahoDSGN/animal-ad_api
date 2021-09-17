package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Push(c *fiber.Ctx) error {
	data := h.services.Parser.Push()

	return c.JSON(fiber.Map{
		"data": data,
	})
}
