package handlers

import (
	authHandler "github.com/dmitriitalent/strittenApi/internal/handlers/auth"
	"github.com/dmitriitalent/strittenApi/internal/services"
)

type Handlers struct {
	authHandler.Auth
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		Auth: authHandler.NewAuthHandler(services),
	}
}