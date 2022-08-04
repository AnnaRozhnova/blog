package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


type errorResponse struct {
	Message string `json:"message"`
} 

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	// logs an error message
	logrus.Error(message)
	// AbortWithStatusJSON stops the chain, writes the status code and return a JSON body
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}