package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"petcard/pkg/database"
	"petcard/pkg/models"
	"strconv"
)

// CreateAd @Router /api/ad/create [post]
func (h *Handler) CreateAd(c *gin.Context) {
	var JSONinput models.Ad

	if err := c.BindJSON(&JSONinput); err != nil {
		return
	}

	data, err := h.services.Ad.Create(JSONinput)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

// GetAllAds @Router /api/ad/ [get]
func (h *Handler) GetAllAds(c *gin.Context) {
	data, err := h.services.Ad.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

// GetAdById @Router /api/ad/:id [get]
func (h *Handler) GetAdById(c *gin.Context) {
	paramId, _ := strconv.Atoi(c.Param("id"))
	if paramId <= 0 {
		return
	}

	data, err := h.services.Ad.GetList(paramId)
	if err != nil {
		return
	}

	if data.Id == 0 {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

// DeleteAd @Router /api/ad/:id [delete]
func (h *Handler) DeleteAd(c *gin.Context) {
	paramId, _ := strconv.Atoi(c.Param("id"))
	if paramId <= 0 {
		return
	}

	data, err := h.services.Ad.Delete(paramId)
	if err != nil {
		return
	}

	if data.Id == 0 {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

// UpdateAd @Router /api/ad/:id [put]
func (h *Handler) UpdateAd(c *gin.Context) {
	var JSONinput models.Ad

	if err := c.BindJSON(&JSONinput); err != nil {
		return
	}

	paramId, _ := strconv.Atoi(c.Param("id"))
	if paramId <= 0 {
		return
	}

	data, err := h.services.Ad.Update(paramId, JSONinput)
	if err != nil {
		return
	}

	if data.Id == 0 {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
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
