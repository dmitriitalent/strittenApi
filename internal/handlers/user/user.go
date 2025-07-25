package userHandler

import (
	"github.com/dmitriitalent/strittenApi/internal/services"
	loggerService "github.com/dmitriitalent/strittenApi/internal/services/logger"
	userService "github.com/dmitriitalent/strittenApi/internal/services/user"
	"github.com/gin-gonic/gin"
)

type User interface {
	GetUser(*gin.Context)
}

type UserHandler struct {
	userService.User
	loggerService.Logger
}

func NewUserHandler(services *services.Services) *UserHandler {
	return &UserHandler{
		User: services.User,
		Logger: services.Logger,
	}
}