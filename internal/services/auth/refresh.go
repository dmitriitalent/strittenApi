package authService

import (
	"fmt"

	jwtService "github.com/dmitriitalent/strittenApi/internal/services/jwt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func (service *AuthService) Refresh(refreshTokenObj *jwt.Token) (accessToken string, refreshToken string, err error) {
	claims, err := service.GetClaims(refreshTokenObj)
	if err != nil {
		return "", "", err
	}

	userId, ok := claims[jwtService.UserIdClaim].(float64)
	if !ok {
		return "", "", fmt.Errorf("cannot convert %s claim to float64", jwtService.UserIdClaim)
	}

	generatedAccessToken, err := service.Jwt.GenerateAccessToken(int(userId), viper.GetString("crypto.accessTokenSalt"))
	if err != nil {
		return "", "", err
	}

	generatedRefreshToken, err := service.Jwt.GenerateRefreshToken(int(userId), viper.GetString("crypto.refreshTokenSalt"))
	if err != nil {
		return "", "", err
	}

	return generatedAccessToken, generatedRefreshToken, nil
}