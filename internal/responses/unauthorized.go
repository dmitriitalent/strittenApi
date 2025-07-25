package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type unauthorizedError struct {
	Message string `json:"message"`
}

func Unauthorized(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedError{message})
}