package eventService

import (
	"fmt"

	"github.com/dmitriitalent/strittenApi/internal/entity"
)

func (service *EventService) CreateEvent(event entity.Event, additionalDatas entity.AdditionalDatas) (entity.Event, entity.AdditionalDatas, error) {
	createdEvent, createdAdditionalDatas, err := service.Repositories.Event.CreateEvent(event, additionalDatas);
	if err != nil {
		return event, additionalDatas, fmt.Errorf("EventService:CreateEvent: %s", err.Error())
	}
	return createdEvent, createdAdditionalDatas, nil
}