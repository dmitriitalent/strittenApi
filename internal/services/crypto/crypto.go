package cryptoService

import loggerService "github.com/dmitriitalent/strittenApi/internal/services/logger"

type Crypto interface {
	HashPassword(password string) (hash string, err error)
	ComparePasswords(password string, hashedPassword string) (err error)
}

type CryptoService struct {
	loggerService.Logger
}

func NewCryptoService(logger loggerService.Logger) *CryptoService {
	return &CryptoService{
		Logger: logger,
	}
}