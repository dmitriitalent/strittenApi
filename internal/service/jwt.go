package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dmitriitalent/strittenApi/internal/repository"
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type JwtService struct {
	repo repository.Jwt
}

func NewJwtService(repo repository.Jwt) *JwtService {
	return &JwtService{repo: repo}
}

func (s *JwtService) GenerateAccessToken(claims TokenClaims, salt string) (string, error) {
	claims.ExpiresAt = time.Now().Add(10 * time.Minute).Unix()
	claims.IssuedAt = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(salt))
}

func (s *JwtService) GenerateRefreshToken(claims TokenClaims, salt string) (string, error) {
	claims.ExpiresAt = time.Now().Add(30 * time.Hour * 24).Unix()
	claims.IssuedAt = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString([]byte(salt))
	if err != nil {
		return "", err
	}

	refreshToken, err = s.addRefreshTokenToDB(claims.UserId, refreshToken)
	if err != nil {
		return "", err
	}

	return refreshToken, err
}

func (s *JwtService) addRefreshTokenToDB(userId int, newRefreshToken string) (string, error) {
	return s.repo.CreateNewRefreshToken(userId, newRefreshToken)
}

func (s *JwtService) ValidateToken(tokenString string, secretKey []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method at JwtService.ValidateToken: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("token validation failed at JwtService.ValidateToken: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token at JwtService.ValidateToken")
	}

	return token, nil
}

func (s *JwtService) GetClaimsFromToken(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims at JwtService.GetClaimsFromToken")
	}

	return claims, nil
}

func (s *JwtService) RemoveRefreshToken(tokenString string) error {
	return s.repo.RemoveRefreshToken(tokenString)
}
