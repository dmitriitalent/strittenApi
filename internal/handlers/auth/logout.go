package authHandler

import (
	"github.com/dmitriitalent/strittenApi/internal/responses"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (handler *AuthHandler) Logout(c *gin.Context) {
	refreshTokenCookie, err := c.Request.Cookie("rt")
	if err != nil {
		responses.Unauthorized(c, err.Error())
		handler.Logger.Debug(err.Error())
		return
	}

	c.SetCookie("at", "", 0, "", "", false, false)
	c.SetCookie("rt", "", 0, "", "", false, false)

	_, err = handler.Jwt.Validate(refreshTokenCookie.Value, []byte(viper.GetString("crypto.refreshTokenSalt")))
	if err != nil {
		responses.Forbidden(c, err.Error())
		return
	}

	if err := handler.Auth.Logout(refreshTokenCookie.Value); err != nil {
		responses.InternalServerError(c)
	}

	responses.Ok(c, "no content")
}