package jwtService

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func (s *JwtService) Validate(tokenString string, secretKey []byte) (*jwt.Token, error) {
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