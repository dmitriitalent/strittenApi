package repository

import (
	"fmt"

	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user entity.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (login, password_hash, name, surname, email) VALUES ($1, $2, $3, $4, $5) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password, user.Name, user.Surname, user.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) FindUser(user entity.User) (int, error) {
	var id int

	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password_hash=$2", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) FindUserById(userId int) (entity.User, error) {
	var user entity.User

	query := fmt.Sprintf("SELECT id, login, name, surname, email FROM %s WHERE id=$1", usersTable)
	row := r.db.QueryRow(query, userId)
	if err := row.Scan(&user.Id, &user.Login, &user.Name, &user.Surname, &user.Email); err != nil {
		return entity.User{}, err
	}

	return user, nil
}
