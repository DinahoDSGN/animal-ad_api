package handler

import "github.com/gofiber/fiber/v2"

// Push @Router /api/parser/push [POST]
func (h *Handler) Push(c *fiber.Ctx) error {
	data := h.services.Parser.Push()

	return c.JSON(fiber.Map{
		"data": data,
	})
}
