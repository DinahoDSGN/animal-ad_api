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
	color := c.Query("color")
	sex := c.Query("sex")
	vaccinated := c.Query("vaccinated")
	spayed := c.Query("spayed")
	passport := c.Query("passport")
	price := c.Query("price")

	values := make(map[string]interface{})

	if specify != "" {
		values["specify"] = specify
	}

	if breed != "" {
		values["breed"] = breed
	}

	if color != "" {
		values["color"] = color
	}

	if sex != "" {
		if sex == "male" {
			values["sex"] = false
		} else if sex == "female" {
			values["sex"] = true
		}
	}

	if vaccinated != "" {
		if vaccinated == "yes" {
			values["vaccinated"] = true
		} else if vaccinated == "no" {
			values["vaccinated"] = false
		}
	}

	if spayed != "" {
		if spayed == "yes" {
			values["spayed"] = true
		} else if spayed == "no" {
			values["spayed"] = false
		}
	}

	if passport != "" {
		if passport == "yes" {
			values["passport"] = true
		} else if passport == "no" {
			values["passport"] = false
		}
	}

	if price != "" {
		values["price"] = price
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
