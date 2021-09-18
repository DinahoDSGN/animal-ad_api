package handler

import (
	"github.com/gofiber/fiber/v2"
	"petcard/pkg/models"
	"strconv"
)

// GetAllUsers @Router /api/user/ [GET]
func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	data, err := h.services.User.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}

// GetUserById @Router /api/user/:id [GET]
func (h *Handler) GetUserById(c *fiber.Ctx) error {
	paramId, _ := strconv.Atoi(c.Params("id"))

	data, err := h.services.User.GetList(paramId)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}

// UpdateUser @Router /api/user/:id [PUT]
func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	var JSONinput models.User

	if err := c.BodyParser(&JSONinput); err != nil {
		return err
	}

	paramId, _ := strconv.Atoi(c.Params("id"))

	data := h.services.User.Update(paramId, JSONinput)

	return c.JSON(fiber.Map{
		"data": data,
	})
}

// DeleteUser @Router /api/user/:id [DELETE]
func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	paramId, _ := strconv.Atoi(c.Params("id"))

	data := h.services.User.Delete(paramId)

	return c.JSON(fiber.Map{
		"data": data,
	})
}
