package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message interface{} `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, status int, message interface{}) error {
	logrus.Error(message)
	c.AbortWithStatusJSON(status, ErrorResponse{message})
	return nil
}
