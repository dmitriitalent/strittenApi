package jwtService

import (
	"github.com/go-faster/errors"
	"github.com/golang-jwt/jwt/v4"
)

func (s *JwtService) GetClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims at JwtService.GetClaims")
	}

	return claims, nil
}

func (s *JwtService) RemoveRefreshToken(tokenString string) error {
	return s.Repositories.RemoveToken(tokenString)
}