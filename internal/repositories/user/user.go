package userRepository

import (
	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/jmoiron/sqlx"
)

const usersTable = "users"

type User interface {
	IsEmailUsed(email string) (err error)
	IsLoginUsed(login string) (err error)
	CreateUser(user entity.User) (id int, err error)
	FindUserByLogin(login string) (user entity.User, err error)
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}
