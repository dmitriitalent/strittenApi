package cryptoService

import (
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func (service *CryptoService) HashPassword(password string) (hash string, err error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), viper.GetInt("crypto.hashpasswordcost"));

	if err != nil {
		service.Logger.Error("Occured while GenerateFromPassword at HashPassword: %w", err)
		return "", err
	}

	return string(hashBytes), nil
}