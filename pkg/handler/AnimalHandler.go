package handler

import (
	"github.com/gofiber/fiber/v2"
	"petcard/pkg/models"
	"strconv"
)

// CreateAnimal @Router /api/spec/create [POST]
func (h *Handler) CreateAnimal(c *fiber.Ctx) error {
	var JSONinput models.Animal

	if err := c.BodyParser(&JSONinput); err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	id, err := h.services.Animal.Create(JSONinput)
	if err != nil {
		return newErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"id":     id,
	})
}

// GetAllAnimals @Router /api/spec/ [GET]
func (h *Handler) GetAllAnimals(c *fiber.Ctx) error {
	data, err := h.services.Animal.GetAll()
	if err != nil {
		return newErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   data,
	})
}

// GetAnimalById @Router /api/spec/:id [GET]
func (h *Handler) GetAnimalById(c *fiber.Ctx) error {
	paramId, _ := strconv.Atoi(c.Params("id"))
	if paramId <= 0 {
		return newErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	data, err := h.services.Animal.GetList(paramId)
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

// DeleteAnimal @Router /api/spec/:id [DELETE]
func (h *Handler) DeleteAnimal(c *fiber.Ctx) error {
	paramId, _ := strconv.Atoi(c.Params("id"))
	if paramId <= 0 {
		return newErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	data, err := h.services.Animal.Delete(paramId)
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

// UpdateAnimal @Router /api/spec/:id [PUT]
func (h *Handler) UpdateAnimal(c *fiber.Ctx) error {
	var JSONinput models.Animal

	if err := c.BodyParser(&JSONinput); err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	paramId, _ := strconv.Atoi(c.Params("id"))
	if paramId <= 0 {
		return newErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	data, err := h.services.Animal.Update(paramId, JSONinput)
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
