package eventService

import (
	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/dmitriitalent/strittenApi/internal/repositories"
)

type Event interface {
	GetEvent(id int) (event entity.Event, err error)
	CreateEvent(entity.Event) (event entity.Event, err error)
}

type EventService struct {
	*repositories.Repositories
}

func NewEventService(repos *repositories.Repositories) *EventService {
	return &EventService{
		Repositories: repos,
	}
}