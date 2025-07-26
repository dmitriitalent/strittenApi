package eventService

import "github.com/dmitriitalent/strittenApi/internal/entity"

func (service *EventService) CreateEvent(event entity.Event) (entity.Event, error) {
	return service.Repositories.Event.CreateEvent(event);
}