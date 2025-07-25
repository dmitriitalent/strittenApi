package validationService

type Validation interface {
	IsEmailValid(email string) (err error)
	IsPasswordValid(password string) (err error)
}

type ValidationService struct {
}

func NewValidationService() *ValidationService {
	return &ValidationService{}
}