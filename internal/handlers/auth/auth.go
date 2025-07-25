package authHandler

import (
	"github.com/dmitriitalent/strittenApi/internal/services"
	authService "github.com/dmitriitalent/strittenApi/internal/services/auth"
	jwtService "github.com/dmitriitalent/strittenApi/internal/services/jwt"
	loggerService "github.com/dmitriitalent/strittenApi/internal/services/logger"
	validationService "github.com/dmitriitalent/strittenApi/internal/services/validation"
	"github.com/gin-gonic/gin"
)

type Auth interface {
	Login(*gin.Context)
	Registration(*gin.Context)
	Refresh(*gin.Context)
	Logout(*gin.Context)
}

type AuthHandler struct {
	authService.Auth
	validationService.Validation
	loggerService.Logger
	jwtService.Jwt
}

func NewAuthHandler(services *services.Services) *AuthHandler {
	return &AuthHandler{
		Auth: services.Auth,
		Jwt: services.Jwt,
		Validation: services.Validation,
		Logger: services.Logger,
	}
}