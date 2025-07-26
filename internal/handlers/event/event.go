package eventHandler

import (
	"github.com/dmitriitalent/strittenApi/internal/services"
	eventService "github.com/dmitriitalent/strittenApi/internal/services/event"
	jwtService "github.com/dmitriitalent/strittenApi/internal/services/jwt"
	loggerService "github.com/dmitriitalent/strittenApi/internal/services/logger"
	"github.com/gin-gonic/gin"
)

type AdditionalDataRowType struct {
	Key 	string `json:"key"`
	Value 	string `json:"value"`
}

type AdditionalDataRowTypeResponse struct {
	Id						int `json:"id"`
	AdditionalDataRowType
	EventId					int `json:"event_id"`
}

type Event interface {
	GetEvent(c *gin.Context)
	CreateEvent(c *gin.Context)
}

type EventHandler struct {
	eventService.Event
	loggerService.Logger
	jwtService.Jwt
}

func NewEventHandler(services *services.Services) *EventHandler {
	return &EventHandler{
		Event: services.Event,
		Logger: services.Logger,
		Jwt: services.Jwt,
	}
}