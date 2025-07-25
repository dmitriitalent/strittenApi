package jwtService

func (service *JwtService) RemoveRefreshTokenFromDB(refreshToken string) (err error) {
	return service.Repositories.RemoveToken(refreshToken)
}