package userRepository

import (
	"fmt"
)

func (r *UserRepository) IsLoginUsed(login string) (err error) {
	var used string = "";

	query := fmt.Sprintf("SELECT login FROM %s WHERE login=$1", usersTable)
	row := r.db.QueryRow(query, login)
	if err := row.Scan(&used); err != nil {
		return  err
	}

	return nil
}