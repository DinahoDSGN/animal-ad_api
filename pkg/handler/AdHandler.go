package handler

import (
	"github.com/gofiber/fiber/v2"
	"petcard/pkg/models"
	"strconv"
)

// CreateAd @Router /api/ad/create [post]
func (h *Handler) CreateAd(c *fiber.Ctx) error {
	var JSONinput models.Ad

	if err := c.BodyParser(&JSONinput); err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	data, err := h.services.Ad.Create(JSONinput)
	if err != nil {
		return newErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   data,
	})
}

// GetAllAds @Router /api/ad/ [get]
func (h *Handler) GetAllAds(c *fiber.Ctx) error {
	data, err := h.services.Ad.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}

// GetAdById @Router /api/ad/:id [get]
func (h *Handler) GetAdById(c *fiber.Ctx) error {
	paramId, _ := strconv.Atoi(c.Params("id"))

	data, err := h.services.Ad.GetList(paramId)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}

// DeleteAd @Router /api/ad/:id [delete]
func (h *Handler) DeleteAd(c *fiber.Ctx) error {
	paramId, _ := strconv.Atoi(c.Params("id"))

	data := h.services.Ad.Delete(paramId)

	return c.JSON(fiber.Map{
		"data": data,
	})
}

// UpdateAd @Router /api/ad/:id [put]
func (h *Handler) UpdateAd(c *fiber.Ctx) error {
	var JSONinput models.Ad

	if err := c.BodyParser(&JSONinput); err != nil {
		return err
	}

	paramId, _ := strconv.Atoi(c.Params("id"))

	data := h.services.Ad.Update(paramId, JSONinput)

	return c.JSON(fiber.Map{
		"data": data,
	})
}
