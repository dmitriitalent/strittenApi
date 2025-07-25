package responses

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func InternalServerError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, time.Now().Unix())
}