package jwtService

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (s *JwtService) GenerateAccessToken(userId int, salt string) (accessToken string, err error) {
	claims := TokenClaims{}
	claims.UserId = userId
	claims.ExpiresAt = time.Now().Add(10 * time.Minute).Unix()
	claims.IssuedAt = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(salt))
}