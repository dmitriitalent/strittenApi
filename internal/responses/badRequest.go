package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type badRequestError struct {
	Message string `json:"message"`
}

func BadRequest(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, badRequestError{message})
}