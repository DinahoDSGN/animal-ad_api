package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//SortAdBy @Router /api/adv/sort [get]
func (h *Handler) SortAdBy(c *gin.Context) {
	specify := c.Query("specify")
	breed := c.Query("breed")
	sex := c.Query("sex")
	values := make(map[string]interface{})
	values["specify"] = ""
	values["breed"] = ""
	values["sex"] = 1

	if specify != "" {
		values["specify"] = specify
	}

	if breed != "" {
		values["breed"] = breed
	}

	if sex != "" {
		if sex == "male" {
			values["sex"] = false
		} else if sex == "female" {
			values["sex"] = true
		}
	}

	data, err := h.services.AdSorts.SortBy(values)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(data) == 0 {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("No matching entries found, "+
			"values: %v", values))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}
