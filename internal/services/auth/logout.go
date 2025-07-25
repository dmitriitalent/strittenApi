package authService

func (service *AuthService) Logout(refreshToken string) error {
	err := service.Jwt.RemoveRefreshTokenFromDB(refreshToken)
	return err
}