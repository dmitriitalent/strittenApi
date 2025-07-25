package handlers

import (
	authHandler "github.com/dmitriitalent/strittenApi/internal/handlers/auth"
	userHandler "github.com/dmitriitalent/strittenApi/internal/handlers/user"
	"github.com/dmitriitalent/strittenApi/internal/services"
)

type Handlers struct {
	authHandler.Auth
	userHandler.User
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		Auth: authHandler.NewAuthHandler(services),
		User: userHandler.NewUserHandler(services),
	}
}