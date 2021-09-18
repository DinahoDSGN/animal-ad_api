package handler

import (
	"github.com/gofiber/fiber/v2"
	"petcard/pkg/models"
	"strconv"
)

// CreateBreed @Router /api/breed/create [POST]
func (h *Handler) CreateBreed(c *fiber.Ctx) error {
	JSONinput := models.Breed{}

	if err := c.BodyParser(&JSONinput); err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	id, err := h.services.Breed.Create(JSONinput)
	if err != nil {
		return newErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"id":     id,
	})
}

// GetAllBreeds @Router /api/breed/ [get]
func (h *Handler) GetAllBreeds(c *fiber.Ctx) error {
	data, err := h.services.Breed.GetAll()
	if err != nil {
		return newErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   data,
	})
}

// GetBreedById @Router /api/breed/:id [get]
func (h *Handler) GetBreedById(c *fiber.Ctx) error {
	paramId, _ := strconv.Atoi(c.Params("id"))
	if paramId <= 0 {
		return newErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	data, err := h.services.Breed.GetList(paramId)
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

// DeleteBreed @Router /api/breed/:id [DELETE]
func (h *Handler) DeleteBreed(c *fiber.Ctx) error {
	paramId, _ := strconv.Atoi(c.Params("id"))
	if paramId <= 0 {
		return newErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	data, err := h.services.Breed.Delete(paramId)
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

// UpdateBreed @Router /api/breed/:id [PUT]
func (h *Handler) UpdateBreed(c *fiber.Ctx) error {
	JSONinput := models.Breed{}

	if err := c.BodyParser(&JSONinput); err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	paramId, _ := strconv.Atoi(c.Params("id"))
	if paramId <= 0 {
		return newErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	data, err := h.services.Breed.Update(paramId, JSONinput)
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
