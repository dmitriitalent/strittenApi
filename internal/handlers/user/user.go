package userHandler

import (
	"github.com/dmitriitalent/strittenApi/internal/services"
	authService "github.com/dmitriitalent/strittenApi/internal/services/auth"
	cryptoService "github.com/dmitriitalent/strittenApi/internal/services/crypto"
	jwtService "github.com/dmitriitalent/strittenApi/internal/services/jwt"
	loggerService "github.com/dmitriitalent/strittenApi/internal/services/logger"
	userService "github.com/dmitriitalent/strittenApi/internal/services/user"
	validationService "github.com/dmitriitalent/strittenApi/internal/services/validation"
	"github.com/gin-gonic/gin"
)

type User interface {
	GetUser(*gin.Context)
	UpdateUser(*gin.Context)
}

type UserHandler struct {
	userService.User
	loggerService.Logger
	jwtService.Jwt
	authService.Auth
	cryptoService.Crypto
	validationService.Validation
}

func NewUserHandler(services *services.Services) *UserHandler {
	return &UserHandler{
		User: services.User,
		Logger: services.Logger,
		Jwt: services.Jwt,
		Auth: services.Auth,
		Crypto: services.Crypto,
		Validation: services.Validation,
	}
}