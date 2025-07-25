package userRepository

import (
	"fmt"

	"github.com/dmitriitalent/strittenApi/internal/entity"
)

func (r *UserRepository) FindUserByLogin(login string) (user entity.User, err error) {
	var result entity.User;

	query := fmt.Sprintf("SELECT id, login, password_hash, name, surname, email FROM %s WHERE login=$1", usersTable)
	row := r.db.QueryRow(query, login)
	if err := row.Scan(&result.Id, &result.Login, &result.Password, &result.Name, &result.Surname, &result.Email); err != nil {
		return entity.User{}, err
	}

	return result, nil
}