package userRepository

import (
	"fmt"

	"github.com/dmitriitalent/strittenApi/internal/entity"
)

func (r *UserRepository) FindUserById(userId int) (entity.User, error) {
	var result entity.User;

	query := fmt.Sprintf("SELECT id, login, password_hash, name, surname, email FROM %s WHERE id=$1", usersTable)
	row := r.db.QueryRow(query, userId)
	if err := row.Scan(&result.Id, &result.Login, &result.Password, &result.Name, &result.Surname, &result.Email); err != nil {
		return entity.User{}, err
	}

	return result, nil
}