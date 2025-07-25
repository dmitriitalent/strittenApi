package services

import (
	"github.com/dmitriitalent/strittenApi/internal/repositories"
	authService "github.com/dmitriitalent/strittenApi/internal/services/auth"
	cryptoService "github.com/dmitriitalent/strittenApi/internal/services/crypto"
	jwtService "github.com/dmitriitalent/strittenApi/internal/services/jwt"
	loggerService "github.com/dmitriitalent/strittenApi/internal/services/logger"
	validationService "github.com/dmitriitalent/strittenApi/internal/services/validation"
)

type Services struct {
	authService.Auth
	validationService.Validation
	loggerService.Logger
	cryptoService.Crypto
	jwtService.Jwt
}

func NewServices(repos *repositories.Repositories) *Services {
	loggerService := loggerService.NewLoggerService()
	jwtService := jwtService.NewJwtService(repos)
	cryptoService := cryptoService.NewCryptoService(loggerService)

	return &Services{
		Auth: authService.NewAuthService(repos, loggerService, jwtService, cryptoService),
		Validation: validationService.NewValidationService(),
		Crypto: cryptoService,
		Logger: loggerService,
		Jwt: jwtService,
	}
}