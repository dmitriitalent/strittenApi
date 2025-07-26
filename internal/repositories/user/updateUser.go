package userRepository

import (
	"errors"
	"fmt"

	"github.com/dmitriitalent/strittenApi/internal/entity"
)

func (r *UserRepository) UpdateUser(user entity.User) (entity.User, error) {
	var updatedUser entity.User;

	tx, err := r.db.Begin()
	if err != nil {
		return user, errors.New("UserRepository:UpdateUser failed to begin transaction: " + err.Error())
	}
	defer tx.Rollback()
	query := fmt.Sprintf("UPDATE %s SET login = $2, name = $3, surname = $4, email = $5, password_hash = $6 WHERE id = $1", usersTable)

	_, err = tx.Exec(query, user.Id, user.Login, user.Name, user.Surname, user.Email, user.Password);
	if err != nil {
		return user, err;
	}

	query = fmt.Sprintf("SELECT id, login, name, surname, email, password_hash FROM %s WHERE id = $1", usersTable)
	queryRow := tx.QueryRow(query, user.Id)
	if err := queryRow.Scan(&updatedUser.Id, &updatedUser.Login, &updatedUser.Name, &updatedUser.Surname, &updatedUser.Email, &updatedUser.Password ); err != nil {
		return user, err;
	}

	if err := tx.Commit(); err != nil {
		return user, errors.New("UserRepository:UpdateUser failed to commit transaction: " + err.Error())
	}

	return updatedUser, nil;
}