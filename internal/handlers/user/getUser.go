package userHandler

import (
	"strconv"

	"github.com/dmitriitalent/strittenApi/internal/responses"
	"github.com/gin-gonic/gin"
)

type GetUserResponse struct {
	Id int `json:"id"`
	Login string `json:"login"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
}

func (handler *UserHandler) GetUser(c *gin.Context) {
	userIdParam := c.Query("user_id")
	loginParam := c.Query("login")
	if userIdParam == "" && loginParam == "" {
		responses.BadRequest(c, "Add user_id or login param")
		handler.Logger.Debug(`user_id and login are empty, "%s", "%s"`, userIdParam, loginParam)
		return
	}

	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		responses.BadRequest(c, "user_id must be number")
		handler.Logger.Debug(`Cannot convert userIdParam (string) to userId int. userIdParam: "%s"`, userIdParam)
		return
	}

	user, err := handler.User.GetUserById(userId)
	if err != nil {
		responses.InternalServerError(c)
		handler.Logger.Debug(err.Error())
		return
	}

	responses.Ok(c, GetUserResponse{
		Id: user.Id,
		Login: user.Login,
		Name: user.Name,
		Surname: user.Surname,
		Email: user.Email,
	})
}