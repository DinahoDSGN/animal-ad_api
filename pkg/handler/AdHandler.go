package handler

import (
	"github.com/gofiber/fiber/v2"
	"petcard/pkg/database"
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
		return newErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   data,
	})
}

// GetAdById @Router /api/ad/:id [get]
func (h *Handler) GetAdById(c *fiber.Ctx) error {
	paramId, _ := strconv.Atoi(c.Params("id"))
	if paramId <= 0 {
		return newErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	data, err := h.services.Ad.GetList(paramId)
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

// DeleteAd @Router /api/ad/:id [delete]
func (h *Handler) DeleteAd(c *fiber.Ctx) error {
	paramId, _ := strconv.Atoi(c.Params("id"))
	if paramId <= 0 {
		return newErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	data, err := h.services.Ad.Delete(paramId)
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

// UpdateAd @Router /api/ad/:id [put]
func (h *Handler) UpdateAd(c *fiber.Ctx) error {
	var JSONinput models.Ad

	if err := c.BodyParser(&JSONinput); err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	paramId, _ := strconv.Atoi(c.Params("id"))
	if paramId <= 0 {
		return newErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	data, err := h.services.Ad.Update(paramId, JSONinput)
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

func (h *Handler) SortBy(c *fiber.Ctx) error {
	paramBy := c.Params("by")
	if paramBy == "asc" {
		data := h.SortByAscending()
		return c.JSON(fiber.Map{
			"status": fiber.StatusOK,
			"method": "Ad price by ascending",
			"data":   data,
		})
	}
	if paramBy == "desc" {
		data := h.SortByDescending()
		return c.JSON(fiber.Map{
			"status": fiber.StatusOK,
			"method": "Ad price by descending",
			"data":   data,
		})
	}

	return nil
}

func (h *Handler) SortByAscending() []models.Ad {
	db := database.DB

	data, _ := h.services.Ad.GetAll()
	db.Model(data).Preload("Animal").Select("ads.*", "animals.*").
		Joins("INNER JOIN animals ON animals.id = ads.animal_id").
		Order("price asc").
		Find(&data)

	return data
}

func (h *Handler) SortByDescending() []models.Ad {
	db := database.DB

	data, _ := h.services.Ad.GetAll()
	db.Model(data).Preload("Animal").Select("ads.*", "animals.*").
		Joins("INNER JOIN animals ON animals.id = ads.animal_id").
		Order("price desc").
		Find(&data)

	return data
}

//func (h *Handler) SortBy(c *fiber.Ctx) error{
//	db := database.DB
//
//	paramType := c.Query("type")
//	if paramType == ""{
//		paramType = "cat"
//	}
//	paramPriceGT := c.Query("price_gt")
//	if paramPriceGT == ""{
//		paramPriceGT = "0"
//	}
//	paramPriceLT := c.Query("price_lt")
//	if paramPriceLT == ""{
//		paramPriceLT = "0"
//	}
//
//	data, _ := h.services.Ad.GetAll()
//
//	db.Model(data).Preload("Animal").
//		Select("ads.*", "animals.*").
//		Joins("INNER JOIN animals ON animals.id = ads.animal_id").
//		Where("type = ?", paramType).
//		Where("price > ?", paramPriceGT).
//		Where("price < ?", paramPriceLT).
//		Find(&data)
//
//	return c.JSON(fiber.Map{
//		"status" : fiber.StatusOK,
//		"method" : "Sort by type",
//		"data" : data,
//	})
//}
