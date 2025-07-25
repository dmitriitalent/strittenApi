package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type forbiddenError struct {
	Message string `json:"message"`
}

func Forbidden(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusForbidden, forbiddenError{message})
}