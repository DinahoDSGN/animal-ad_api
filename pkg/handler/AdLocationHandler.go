package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"petcard/pkg/models"
	"strconv"
)

func (h *Handler) CreateAdLocation(c *gin.Context) {
	var JSONInput models.AdLocation

	if err := c.BindJSON(&JSONInput); err != nil {
		return
	}

	data, err := h.services.AdLocation.Create(JSONInput)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

func (h *Handler) GetAllLocations(c *gin.Context) {
	data, err := h.services.AdLocation.GetAll()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

func (h *Handler) GetLocationById(c *gin.Context) {
	paramId, _ := strconv.Atoi(c.Param("id"))
	if paramId <= 0 {
		return
	}

	data, err := h.services.AdLocation.GetList(paramId)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})

}
