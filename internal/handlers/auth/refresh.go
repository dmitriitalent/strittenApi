package authHandler

import (
	"time"

	"github.com/dmitriitalent/strittenApi/internal/responses"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (handler *AuthHandler) Refresh(c *gin.Context) {
	refreshTokenCookie, err := c.Request.Cookie("rt")
	if err != nil {
		responses.Unauthorized(c, err.Error())
		handler.Logger.Debug(err.Error())
		return
	}

	token, err := handler.Jwt.Validate(refreshTokenCookie.Value, []byte(viper.GetString("crypto.refreshTokenSalt")))
	if err != nil {
		responses.Unauthorized(c, err.Error())
		handler.Logger.Debug(err.Error())
		return
	}

	accessToken, refreshToken, err := handler.Auth.Refresh(token)
	if err != nil {
		responses.InternalServerError(c)
		handler.Logger.Debug(err.Error())
		return
	}

	c.SetCookie(
		"at",
		accessToken,
		int(time.Now().Add(10*time.Minute).Unix()),
		"/",
		"",
		true,
		false,
	)

	c.SetCookie(
		"rt",
		refreshToken,
		int(time.Now().Add(30*time.Hour*24).Unix()),
		"/",
		"",
		true,
		true,
	)

	responses.Ok(c, map[string]interface{}{
		"at": accessToken,
	})
}