package repository

import (
	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(user entity.User) (int, error)
	FindUser(user entity.User) (int, error)
	FindUserById(userId int) (entity.User, error)
}

type Jwt interface {
	CreateNewRefreshToken(userId int, newRefreshToken string) (string, error)
	RemoveRefreshToken(tokenString string) error
}

type Event interface {

}

type Repository struct {
	User
	Jwt
	Event
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:  NewUserPostgres(db),
		Jwt:   NewJwtPostgres(db),
		Event: NewEventPostgres(db),
	}
}
