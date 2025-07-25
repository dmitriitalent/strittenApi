package authService

func (service *AuthService) IsLoginUsed(login string) (err error) {
	return service.User.IsLoginUsed(login)
}