package authService

import (
	"errors"

	"github.com/spf13/viper"
)

func (service *AuthService) Login(login string, password string) (accessToken string, refreshToken string, err error) {
	user, err := service.User.FindUserByLogin(login)
	if err != nil {
		if(err.Error() == "sql: no rows in result set") { 
			return "", "", errors.New("User does not exist")
		}
		return "", "", err
	}

	err = service.Crypto.ComparePasswords(password, user.Password)
	if err != nil {
		if(err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password") { 
			return "", "", errors.New("Incorrect login or password")
		}
		return "", "", err
	}

	generatedAccessToken, err := service.GenerateAccessToken(user.Id, viper.GetString("crypto.accessTokenSalt"))
	if err != nil {
		return "", "", err
	}

	generatedRefreshToken, err := service.GenerateRefreshToken(user.Id, viper.GetString("crypto.refreshTokenSalt"))
	if err != nil {
		return "", "", err
	}

	return generatedAccessToken, generatedRefreshToken, nil
}