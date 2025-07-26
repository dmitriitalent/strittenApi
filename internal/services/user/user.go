package userService

import (
	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/dmitriitalent/strittenApi/internal/repositories"
	cryptoService "github.com/dmitriitalent/strittenApi/internal/services/crypto"
)

type User interface {
	GetUserById(userId int) (user entity.User, err error)
	UpdateUser(newUser entity.User) (updatedUser entity.User, err error)
}

type UserService struct {
	*repositories.Repositories
	cryptoService.Crypto
}

func NewUserSerivce(repos *repositories.Repositories, cryptoService cryptoService.Crypto) *UserService {
	return &UserService{
		Repositories: repos,
		Crypto: cryptoService,
	}

}