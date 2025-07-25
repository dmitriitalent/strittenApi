package authService

import (
	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/dmitriitalent/strittenApi/internal/repositories"
	cryptoService "github.com/dmitriitalent/strittenApi/internal/services/crypto"
	jwtService "github.com/dmitriitalent/strittenApi/internal/services/jwt"
	loggerService "github.com/dmitriitalent/strittenApi/internal/services/logger"
	"github.com/golang-jwt/jwt/v4"
)

type Auth interface {
	Login(login string, password string) (accessToken string, refreshToken string, err error)
	Registration(user entity.User) (accessToken string, refreshToken string, err error)
	Refresh(refreshTokenObj *jwt.Token) (accessToken string, refreshToken string, err error)
	Logout(refreshToken string) error
	IsEmailUsed(email string) (err error)
	IsLoginUsed(login string) (err error)
}

type AuthService struct {
	*repositories.Repositories
	jwtService.Jwt
	loggerService.Logger
	cryptoService.Crypto
}

func NewAuthService(repos *repositories.Repositories, logger loggerService.Logger, jwtService jwtService.Jwt, cryptoService cryptoService.Crypto) *AuthService {
	return &AuthService{
		Repositories: repos,
		Logger: logger,
		Jwt: jwtService,
		Crypto: cryptoService,
	}
}