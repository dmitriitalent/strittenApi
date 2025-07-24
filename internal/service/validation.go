package service

import (
	"errors"
	"regexp"
)

type ValidationService struct {
}

func NewValidationService() *ValidationService {
	return &ValidationService{}
}

func (s ValidationService) IsEmailValid(email string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{1,30}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("email is not valid")
	}
	return nil
}

func (s ValidationService) IsPasswordValid(password string) error {
	passwordRegex := regexp.MustCompile(`^[a-z0-9._!@#$%^&*]{7,50}$`)
	if !passwordRegex.MatchString(password) {
		return errors.New("password is not valid")
	}
	return nil
}
