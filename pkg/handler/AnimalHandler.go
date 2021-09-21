package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"petcard/pkg/models"
	"strconv"
)

// CreateAnimal @Router /api/spec/create [POST]
func (h *Handler) CreateAnimal(c *gin.Context) {
	var JSONinput models.Animal

	if err := c.BindJSON(&JSONinput); err != nil {
		return
	}

	id, err := h.services.Animal.Create(JSONinput)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": id,
	})
}

// GetAllAnimals @Router /api/spec/ [GET]
func (h *Handler) GetAllAnimals(c *gin.Context) {
	data, err := h.services.Animal.GetAll()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

// GetAnimalById @Router /api/spec/:id [GET]
func (h *Handler) GetAnimalById(c *gin.Context) {
	paramId, _ := strconv.Atoi(c.Param("id"))
	if paramId <= 0 {
		return
	}

	data, err := h.services.Animal.GetList(paramId)
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

// DeleteAnimal @Router /api/spec/:id [DELETE]
func (h *Handler) DeleteAnimal(c *gin.Context) {
	paramId, _ := strconv.Atoi(c.Param("id"))
	if paramId <= 0 {
		return
	}

	data, err := h.services.Animal.Delete(paramId)
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

// UpdateAnimal @Router /api/spec/:id [PUT]
func (h *Handler) UpdateAnimal(c *gin.Context) {
	var JSONinput models.Animal

	if err := c.BindJSON(&JSONinput); err != nil {
		return
	}

	paramId, _ := strconv.Atoi(c.Param("id"))
	if paramId <= 0 {
		return
	}

	data, err := h.services.Animal.Update(paramId, JSONinput)
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
