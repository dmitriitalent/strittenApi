package eventHandler

import (
	"strconv"
	"time"

	"github.com/dmitriitalent/strittenApi/internal/responses"
	"github.com/gin-gonic/gin"
)

type GetEventResponse struct {
	Id 			int
	Name 		string
	Description string
	Place 		string
	Date 		time.Time
	Count 		int
	Fundraising string
	UserId 		int
} 

func (handler *EventHandler) GetEvent(c *gin.Context) {
	eventIdParam := c.Query("id");
	if eventIdParam == "" {
		responses.BadRequest(c, "Add id or login param")
		handler.Logger.Debug(`id and login are empty, "%s"`, eventIdParam)
		return
	}

	eventId, err := strconv.Atoi(eventIdParam)
	if err != nil {
		responses.BadRequest(c, "id must be number")
		handler.Logger.Debug(`Cannot conver eventIdParam to int: %s`, eventIdParam)
		return
	}

	event, err := handler.Event.GetEvent(eventId)
	if err != nil {
		responses.InternalServerError(c)
		handler.Logger.Debug(err.Error())
		return
	}

	responses.Ok(c, GetEventResponse{
		Id: 			event.Id,
		Name: 			event.Name,
		Description: 	event.Description,
		Place: 			event.Place,
		Date: 			event.Date,
		Count: 			event.Count,
		Fundraising: 	event.Fundraising,
		UserId: 		event.UserId,
	})
}