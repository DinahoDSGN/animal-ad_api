package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"petcard/pkg/models"
	"strconv"
)

// GetAllUsers @Router /api/user/ [GET]
func (h *Handler) GetAllUsers(c *gin.Context) {
	data, err := h.services.User.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

func (h *Handler) GetUser(c *gin.Context) {
	data, err := h.services.User.GetList(h.services.GetUserId())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if data.Id == 0 {
		newErrorResponse(c, http.StatusInternalServerError, "not authorized")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

// GetUserById @Router /api/user/:id [GET]
func (h *Handler) GetUserById(c *gin.Context) {
	paramId, _ := strconv.Atoi(c.Param("id"))
	if paramId <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	data, err := h.services.User.GetList(paramId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if data.Id == 0 {
		newErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

// UpdateUser @Router /api/user/:id [PUT]
func (h *Handler) UpdateUser(c *gin.Context) {
	JSONinput := models.User{}

	if err := c.BindJSON(&JSONinput); err != nil {
		return
	}

	paramId, _ := strconv.Atoi(c.Param("id"))
	if paramId <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	data, err := h.services.User.Update(paramId, JSONinput)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if data.Id == 0 {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

func (h *Handler) UpdateRating(c *gin.Context) {
	JSONinput := models.User{}

	if err := c.BindJSON(&JSONinput); err != nil {
		return
	}

	paramId, _ := strconv.Atoi(c.Param("id"))
	if paramId <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	rating, err := h.services.User.UpdateRating(paramId, JSONinput)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if rating == 0 {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]float32{
		"data": rating,
	})
}

// DeleteUser @Router /api/user/:id [DELETE]
func (h *Handler) DeleteUser(c *gin.Context) {
	paramId, _ := strconv.Atoi(c.Param("id"))
	if paramId <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	data, err := h.services.User.Delete(paramId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if data.Id == 0 {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}
