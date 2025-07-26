package eventRepository

import (
	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/jmoiron/sqlx"
)

const eventsTable = "events"

type Event interface {
	GetEvent(id int) (event entity.Event, err error)
	CreateEvent(event entity.Event, additionalDatas entity.AdditionalDatas) (createdEvent entity.Event, createdAdditionalDatas entity.AdditionalDatas, err error)
}

type EventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}