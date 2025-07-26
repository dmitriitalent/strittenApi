package userHandler

import (
	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/dmitriitalent/strittenApi/internal/responses"
	jwtService "github.com/dmitriitalent/strittenApi/internal/services/jwt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type UpdateUserRequest struct {
	Login 	 string `json:"login"`
	Name 	 string `json:"name"`
	Surname  string `json:"surname"`
	Email 	 string `json:"email"`
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
} 

type UpdateUserResponse struct {
	Id	 	 int 	`json:"id"`
	Login 	 string `json:"login"`
	Name 	 string `json:"name"`
	Surname  string `json:"surname"`
	Email 	 string `json:"email"`
} 

func (handler *UserHandler) UpdateUser(c *gin.Context) {
	var dto UpdateUserRequest;
	
	err := c.BindJSON(&dto);
	if err != nil {
		responses.BadRequest(c, "Some field not found")
		handler.Logger.Debug(err.Error())
		return
	}

	refreshToken, err := c.Cookie("rt")
	if err != nil {
		responses.BadRequest(c, "rt (refresh token) is empty or not found")
		handler.Logger.Debug(err.Error())
		return
	}

	refreshTokenObj, err := handler.Jwt.Validate(refreshToken, []byte(viper.GetString("crypto.refreshTokenSalt")))
	if err != nil {
		responses.Unauthorized(c, "rt (refresh token) is not valid")
		handler.Logger.Debug(err.Error())
		return
	}

	claims, err := handler.Jwt.GetClaims(refreshTokenObj);
	if err != nil {
		responses.Unauthorized(c, "Invalid claims")
		handler.Logger.Debug(err.Error())
		return
	}

	userId := int(claims[jwtService.UserIdClaim].(float64))

	oldUser, err := handler.User.GetUserById(userId);
	if err != nil {
		responses.InternalServerError(c)
		handler.Logger.Debug("UserHandler:UpdateUser:GetUserById: %s", err.Error())
		return
	}

	err = handler.Auth.IsEmailUsed(dto.Email)
	if(err == nil && oldUser.Email != dto.Email) {
		responses.BadRequest(c, "Email is already used")
		handler.Logger.Debug("Email is already used")
		return
	}

	err = handler.Auth.IsLoginUsed(dto.Login)
	if(err == nil && oldUser.Login != dto.Login) {
		responses.BadRequest(c, "Login is already used")
		handler.Logger.Debug("Login is already used")
		return
	}

	if err := handler.Crypto.ComparePasswords(dto.OldPassword, oldUser.Password); err != nil {
		responses.Forbidden(c, "Invalid password")
		handler.Logger.Debug("Password entered to confirm authorization and the old password do not match: %s", err.Error())
		return
	}

	if err := handler.Validation.IsEmailValid(dto.Email); err != nil {
		responses.BadRequest(c, "Email is not valid")
		handler.Logger.Debug("Email is not valid: %s", dto.Email)
		return
	}

	if err := handler.Validation.IsPasswordValid(dto.NewPassword); err != nil {
		responses.BadRequest(c, "Password is not valid")
		handler.Logger.Debug("Password is not valid: %s", dto.NewPassword)
		return
	}

	var newUser = entity.User{
		Id: 		userId,
		Login: 		dto.Login,
		Name: 		dto.Name,
		Surname: 	dto.Surname,
		Email: 		dto.Email,
		Password: 	dto.NewPassword,
	}

	updatedUser, err := handler.User.UpdateUser(newUser);
	if err != nil {
		responses.InternalServerError(c)
		handler.Logger.Debug(err.Error())
		return
	}

	responses.Ok(c, UpdateUserResponse{
		Id: 		updatedUser.Id,
		Login: 		updatedUser.Login,
		Name: 		updatedUser.Name,
		Surname: 	updatedUser.Surname,
		Email: 		updatedUser.Email,
	})
}