package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Push @Router /api/parser/push [POST]
func (h *Handler) Push(c *gin.Context) {
	data := h.services.Parser.Push()

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}
