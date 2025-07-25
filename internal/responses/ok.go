package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context, payload interface{}) {
	c.JSON(http.StatusOK, payload);
}
