package authHandler

import (
	"time"

	"github.com/dmitriitalent/strittenApi/internal/responses"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

func (handler *AuthHandler) Login(c *gin.Context) {
	var dto LoginRequest; 
	
	if err := c.BindJSON(&dto); err != nil {
		responses.BadRequest(c, "Fields login or password not found")
		handler.Logger.Debug(err.Error())
		return
	}

	accessToken, refreshToken, err := handler.Auth.Login(dto.Login, dto.Password);
	if(err != nil) {
		switch (err.Error()) {
			case "User does not exist":
				responses.BadRequest(c, "User does not exist")
				handler.Logger.Debug(err.Error())
				return 
			case "Incorrect login or password":
				responses.BadRequest(c, "Incorrect login or password")
				handler.Logger.Debug(err.Error())
				return 
			default: 
				responses.InternalServerError(c)
				handler.Logger.Debug(err.Error())
				return 
		}
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