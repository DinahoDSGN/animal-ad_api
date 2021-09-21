package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"petcard/pkg/models"
)

type SignInInput struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

func (h *Handler) SignIn(c *gin.Context) {
	var JSONinput SignInInput

	if err := c.BindJSON(&JSONinput); err != nil {
		return
	}

	token, err := h.services.Authorization.GenerateToken(JSONinput.Username, JSONinput.Password)
	if err != nil {
		newErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	if token == "" {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) SignUp(c *gin.Context) {
	var JSONinput models.User

	if err := c.BindJSON(&JSONinput); err != nil {
		return
	}

	data, err := h.services.Authorization.SignUp(JSONinput)
	if err != nil {
		newErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}
