package validationService

import (
	"errors"
	"regexp"
)

func (s *ValidationService) IsEmailValid(email string) (err error) {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{1,30}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("email is not valid")
	}
	return nil
}