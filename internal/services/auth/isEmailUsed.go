package authService

func (service *AuthService) IsEmailUsed(email string) (err error) {
	return service.User.IsEmailUsed(email)
}