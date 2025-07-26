package handlers

import (
	authHandler "github.com/dmitriitalent/strittenApi/internal/handlers/auth"
	eventHandler "github.com/dmitriitalent/strittenApi/internal/handlers/event"
	userHandler "github.com/dmitriitalent/strittenApi/internal/handlers/user"
	"github.com/dmitriitalent/strittenApi/internal/services"
)

type Handlers struct {
	authHandler.Auth
	userHandler.User
	eventHandler.Event
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		Auth: authHandler.NewAuthHandler(services),
		User: userHandler.NewUserHandler(services),
		Event: eventHandler.NewEventHandler(services),
	}
}