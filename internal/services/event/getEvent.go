package eventService

import (
	"fmt"

	"github.com/dmitriitalent/strittenApi/internal/entity"
)

func (service EventService) GetEvent(id int) (entity.Event, entity.AdditionalDatas, error) {
	event, additionalData, err := service.Repositories.Event.GetEvent(id);
	if err != nil {
		return entity.Event{}, entity.AdditionalDatas{}, fmt.Errorf("EventService:GetEvent: %s", err.Error())
	}

	return event, additionalData, nil 
}