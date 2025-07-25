package userService

import (
	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/dmitriitalent/strittenApi/internal/repositories"
)

type User interface {
	GetUserById(userId int) (user entity.User, err error)
}

type UserService struct {
	*repositories.Repositories
}

func NewUserSerivce(repos *repositories.Repositories) *UserService {
	return &UserService{
		Repositories: repos,
	}
}