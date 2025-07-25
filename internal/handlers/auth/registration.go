package authHandler

import (
	"time"

	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/dmitriitalent/strittenApi/internal/responses"
	"github.com/gin-gonic/gin"
)

type RegistrationRequest struct {
	Login string `json:"login"`
	Password string `json:"password"`
	Email string `json:"email"`
	Name string `json:"name"`
	Surname string `json:"surname"`
}

func (handler *AuthHandler) Registration(c *gin.Context) {
	var dto RegistrationRequest; 
	if err := c.BindJSON(&dto); err != nil {
		responses.BadRequest(c, "Some fields not found")
		handler.Logger.Debug(err.Error())
		return
	}

	user := entity.User{
		Login: dto.Login,
		Password: dto.Password,
		Email: dto.Email,
		Name: dto.Name,
		Surname: dto.Surname,
	}

	err := handler.Auth.IsEmailUsed(user.Email)
	if(err == nil) {
		responses.BadRequest(c, "Email is already used")
		handler.Logger.Debug("Email is already used")
		return
	}

	err = handler.Auth.IsLoginUsed(user.Login)
	if(err == nil) {
		responses.BadRequest(c, "Login is already used")
		handler.Logger.Debug("Login is already used")
		return
	}

	accessToken, refreshToken, err := handler.Auth.Registration(user);
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