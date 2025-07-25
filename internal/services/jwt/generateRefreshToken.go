package jwtService

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (s *JwtService) GenerateRefreshToken(userId int, salt string) (string, error) {
	claims := TokenClaims{}
	claims.UserId = userId
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