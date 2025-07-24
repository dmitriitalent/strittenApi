package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/dmitriitalent/strittenApi/internal/repository"
)

type User interface {
	CreateUser(user entity.User) (int, error)
	FindUser(user entity.User) (int, error)
	FindUserById(userId int) (entity.User, error)
}

type Validation interface {
	IsPasswordValid(password string) error
	IsEmailValid(email string) error
}

type Jwt interface {
	GenerateAccessToken(claims TokenClaims, AccessTokenSalt string) (string, error)
	GenerateRefreshToken(claims TokenClaims, RefreshTokenSalt string) (string, error)
	ValidateToken(tokenString string, secretKey []byte) (*jwt.Token, error)
	GetClaimsFromToken(token *jwt.Token) (jwt.MapClaims, error)
	RemoveRefreshToken(tokenString string) error
}

type Event interface {
	CreateEvent() error;
}

type Service struct {
	User
	Jwt
	Validation
	Event
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:       NewUserService(repo.User),
		Jwt:        NewJwtService(repo.Jwt),
		Validation: NewValidationService(),
		Event: 		NewEventService(repo.Event),
	}
}
