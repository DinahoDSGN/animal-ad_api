package handler

import (
	"github.com/gofiber/fiber/v2"
	"petcard/pkg/models"
	"strconv"
)

// CreateSpec @Router /api/spec/create [post]
func (h *Handler) CreateSpec(c *fiber.Ctx) error {
	var JSONinput models.Specify

	if err := c.BodyParser(&JSONinput); err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	id, err := h.services.Spec.Create(JSONinput)
	if err != nil {
		return newErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"id":     id,
	})
}

// GetAllSpecs @Router /api/spec/ [get]
func (h *Handler) GetAllSpecs(c *fiber.Ctx) error {
	data, err := h.services.Spec.GetAll()
	if err != nil {
		return newErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   data,
	})
}

// GetSpecById @Router /api/spec/:id [get]
func (h *Handler) GetSpecById(c *fiber.Ctx) error {
	paramId, _ := strconv.Atoi(c.Params("id"))
	if paramId <= 0 {
		return newErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	data, err := h.services.Spec.GetList(paramId)
	if err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if data.Id == 0 {
		return newErrorResponse(c, fiber.StatusBadRequest, "record not found")
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   data,
	})
}

// DeleteSpec @Router /api/spec/:id [delete]
func (h *Handler) DeleteSpec(c *fiber.Ctx) error {
	paramId, _ := strconv.Atoi(c.Params("id"))
	if paramId <= 0 {
		return newErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	data, err := h.services.Spec.Delete(paramId)
	if err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if data.Id == 0 {
		return newErrorResponse(c, fiber.StatusBadRequest, "record not found")
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
	})
}

// UpdateSpec @Router /api/spec/:id [put]
func (h *Handler) UpdateSpec(c *fiber.Ctx) error {
	var JSONinput models.Specify

	if err := c.BodyParser(&JSONinput); err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	paramId, _ := strconv.Atoi(c.Params("id"))
	if paramId <= 0 {
		return newErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	data, err := h.services.Spec.Update(paramId, JSONinput)
	if err != nil {
		return newErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	if data.Id == 0 {
		return newErrorResponse(c, fiber.StatusBadRequest, "record not found")
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   data,
	})
}
