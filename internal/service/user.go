package service

import (
	"crypto/sha256"
	"fmt"

	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/dmitriitalent/strittenApi/internal/repository"
)

const salt = "123"

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user entity.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *UserService) FindUser(user entity.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.FindUser(user)
}

func (s *UserService) generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *UserService) FindUserById(userId int) (entity.User, error) {
	return s.repo.FindUserById(userId)
}
