package cryptoService

import "golang.org/x/crypto/bcrypt"

func (service *CryptoService) ComparePasswords(password string, hashedPassword string) (err error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err;
	}
	
	return nil;
}