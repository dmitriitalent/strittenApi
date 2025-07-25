package jwtService

func (s *JwtService) addRefreshTokenToDB(userId int, newToken string) (string, error) {
	return s.Repositories.CreateToken(userId, newToken)
}