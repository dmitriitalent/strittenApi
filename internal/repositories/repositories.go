package repositories

import (
	refreshTokenRepository "github.com/dmitriitalent/strittenApi/internal/repositories/refreshToken"
	userRepository "github.com/dmitriitalent/strittenApi/internal/repositories/user"
	"github.com/jmoiron/sqlx"
)


type Repositories struct {
	refreshTokenRepository.RefreshToken
	userRepository.User
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		RefreshToken: refreshTokenRepository.NewRefreshTokenRepository(db),
		User: userRepository.NewUserRepository(db),
	}
}
