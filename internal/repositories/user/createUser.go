package userRepository

import (
	"fmt"

	"github.com/dmitriitalent/strittenApi/internal/entity"
)

func (r *UserRepository) CreateUser(user entity.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (login, password_hash, name, surname, email) VALUES ($1, $2, $3, $4, $5) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password, user.Name, user.Surname, user.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}