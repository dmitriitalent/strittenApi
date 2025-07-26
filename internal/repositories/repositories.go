package repositories

import (
	eventRepository "github.com/dmitriitalent/strittenApi/internal/repositories/event"
	refreshTokenRepository "github.com/dmitriitalent/strittenApi/internal/repositories/refreshToken"
	userRepository "github.com/dmitriitalent/strittenApi/internal/repositories/user"
	"github.com/jmoiron/sqlx"
)


type Repositories struct {
	refreshTokenRepository.RefreshToken
	userRepository.User
	eventRepository.Event
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		RefreshToken: refreshTokenRepository.NewRefreshTokenRepository(db),
		User: userRepository.NewUserRepository(db),
		Event: eventRepository.NewEventRepository(db),
	}
}
