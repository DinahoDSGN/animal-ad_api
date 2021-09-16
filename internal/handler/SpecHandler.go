package handler

import (
	"github.com/gofiber/fiber/v2"
	"petcard/internal/models"
	"strconv"
)

// CreateSpec @Router /api/spec/create [post]
func (h *Handler) CreateSpec(c *fiber.Ctx) error {
	var JSONinput models.Specify

	if err := c.BodyParser(&JSONinput); err != nil {
		return err
	}

	data := h.services.Spec.Create(JSONinput)

	return c.JSON(fiber.Map{
		"data": data,
	})
}

// GetAllSpecs @Router /api/spec/ [get]
func (h *Handler) GetAllSpecs(c *fiber.Ctx) error {
	data, err := h.services.Spec.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}

// GetSpecById @Router /api/spec/:id [get]
func (h *Handler) GetSpecById(c *fiber.Ctx) error {
	paramId, _ := strconv.Atoi(c.Params("id"))

	data, err := h.services.Spec.GetList(paramId)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}

// DeleteSpec @Router /api/spec/:id [delete]
func (h *Handler) DeleteSpec(c *fiber.Ctx) error {
	paramId, _ := strconv.Atoi(c.Params("id"))

	data, err := h.services.Spec.GetList(paramId)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}

// UpdateSpec @Router /api/spec/:id [put]
func (h *Handler) UpdateSpec(c *fiber.Ctx) error {
	var JSONinput models.Specify

	if err := c.BodyParser(&JSONinput); err != nil {
		return err
	}

	paramId, _ := strconv.Atoi(c.Params("id"))

	data := h.services.Spec.Update(paramId, JSONinput)

	return c.JSON(fiber.Map{
		"data": data,
	})
}
