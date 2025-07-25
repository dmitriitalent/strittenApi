package validationService

import (
	"errors"
	"regexp"
)

func (s *ValidationService) IsPasswordValid(password string) (err error) {
	passwordRegex := regexp.MustCompile(`^[a-z0-9._!@#$%^&*]{7,50}$`)
	if !passwordRegex.MatchString(password) {
		return errors.New("password is not valid")
	}
	return nil
}
