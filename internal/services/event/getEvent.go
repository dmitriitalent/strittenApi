package eventService

import "github.com/dmitriitalent/strittenApi/internal/entity"

func (service EventService) GetEvent(id int) (entity.Event, error) {
	return service.Repositories.Event.GetEvent(id);
}