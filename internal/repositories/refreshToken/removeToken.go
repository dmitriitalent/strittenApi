package refreshTokenRepository

import "fmt"

func (r *RefreshTokenRepository) RemoveToken(tokenString string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE refresh_token=$1", refreshTokensTable)
	_, err := r.db.Exec(query, tokenString)
	if err != nil {
		return err
	}

	return nil
}
