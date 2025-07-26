package eventRepository

import (
	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/jmoiron/sqlx"
)

const eventsTable = "events"

type Event interface {
	GetEvent(id int) (event entity.Event, err error)
	CreateEvent(event entity.Event) (createdEvent entity.Event, err error)
}

type EventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}