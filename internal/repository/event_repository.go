package repository

import (
	"fmt"

	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/jmoiron/sqlx"
)

type EventPostgres struct {
	db *sqlx.DB
}

func NewEventPostgres(db *sqlx.DB) *EventPostgres {
	return &EventPostgres{db: db}
}

func (r *EventPostgres) CreateEvent(event entity.Event) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, description, place, date, count, fundraising) VALUES ($1, $2, $3, $4, $5) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password, user.Name, user.Surname, user.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}


