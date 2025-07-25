package userRepository

import (
	"fmt"
)

func (r *UserRepository) IsEmailUsed(email string) (err error) {
	var used string = "";

	query := fmt.Sprintf("SELECT email FROM %s WHERE email=$1", usersTable)
	row := r.db.QueryRow(query, email)
	if err := row.Scan(&used); err != nil {
		return err
	}

	return nil
}