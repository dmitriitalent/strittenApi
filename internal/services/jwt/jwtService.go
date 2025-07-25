package jwtService

import (
	"github.com/dmitriitalent/strittenApi/internal/repositories"
	"github.com/golang-jwt/jwt/v4"
)

const UserIdClaim = "user_id"

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type Jwt interface {
	addRefreshTokenToDB(userId int, newToken string) (string, error)
	GetClaims(token *jwt.Token) (jwt.MapClaims, error)
	GenerateAccessToken(userId int, salt string) (accessToken string, err error)
	GenerateRefreshToken(userId int, salt string) (refreshToken string, err error)
	Validate(tokenString string, secretKey []byte) (*jwt.Token, error)
	RemoveRefreshTokenFromDB(refreshToken string) (err error)
}

type JwtService struct {
	*repositories.Repositories
}

func NewJwtService(repos *repositories.Repositories) *JwtService {
	return &JwtService{
		Repositories: repos,
	}
}