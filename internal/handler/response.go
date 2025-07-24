package handler

import (
	"github.com/gin-gonic/gin"
)

type error struct {
	Message string `json:"message"`
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, error{message})
}
