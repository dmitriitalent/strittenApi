package authService

import (
	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/spf13/viper"
)

// TODO добавить горутины чтобы токены генерировались парралельно
func (service *AuthService) Registration(user entity.User) (accessToken string, refreshToken string, err error) {
	user.Password, err = service.Crypto.HashPassword(user.Password)
	if err != nil {
		return "", "", err
	}

	id, err := service.Repositories.User.CreateUser(user)
	if err != nil {
		return "", "", err
 	}

	generatedAccessToken, err := service.Jwt.GenerateAccessToken(id, viper.GetString("crypto.accessTokenSalt"))
	if err != nil {
		return "", "", err
	}

	generatedRefreshToken, err := service.Jwt.GenerateRefreshToken(id, viper.GetString("crypto.refreshTokenSalt"))
	if err != nil {
		return "", "", err
	}

	return generatedAccessToken, generatedRefreshToken, nil
}