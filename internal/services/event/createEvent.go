package eventService

import (
	"fmt"

	"github.com/dmitriitalent/strittenApi/internal/entity"
)

func (service *EventService) CreateEvent(event entity.Event, additionalDatas entity.AdditionalDatas) (entity.Event, error) {
	event, err := service.Repositories.Event.CreateEvent(event, additionalDatas);
	return event, fmt.Errorf("EventService:CreateEvent: %s", err.Error())
}