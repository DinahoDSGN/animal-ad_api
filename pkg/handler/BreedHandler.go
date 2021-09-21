package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"petcard/pkg/models"
	"strconv"
)

// CreateBreed @Router /api/breed/create [POST]
func (h *Handler) CreateBreed(c *gin.Context) {
	JSONinput := models.Breed{}

	if err := c.BindJSON(&JSONinput); err != nil {
		return
	}

	id, err := h.services.Breed.Create(JSONinput)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// GetAllBreeds @Router /api/breed/ [get]
func (h *Handler) GetAllBreeds(c *gin.Context) {
	data, err := h.services.Breed.GetAll()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

// GetBreedById @Router /api/breed/:id [get]
func (h *Handler) GetBreedById(c *gin.Context) {
	paramId, _ := strconv.Atoi(c.Param("id"))
	if paramId <= 0 {
		return
	}

	data, err := h.services.Breed.GetList(paramId)
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

// DeleteBreed @Router /api/breed/:id [DELETE]
func (h *Handler) DeleteBreed(c *gin.Context) {
	paramId, _ := strconv.Atoi(c.Param("id"))
	if paramId <= 0 {
		return
	}

	data, err := h.services.Breed.Delete(paramId)
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

// UpdateBreed @Router /api/breed/:id [PUT]
func (h *Handler) UpdateBreed(c *gin.Context) {
	JSONinput := models.Breed{}

	if err := c.BindJSON(&JSONinput); err != nil {
		return
	}

	paramId, _ := strconv.Atoi(c.Param("id"))
	if paramId <= 0 {
		return
	}

	data, err := h.services.Breed.Update(paramId, JSONinput)
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
