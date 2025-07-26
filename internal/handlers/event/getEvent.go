package eventHandler

import (
	"strconv"
	"time"

	"github.com/dmitriitalent/strittenApi/internal/responses"
	"github.com/gin-gonic/gin"
)

type GetEventResponse struct {
	Id 				int								`json:"id"`
	Name 			string							`json:"name"`
	Description 	string							`json:"description"`
	Place 			string							`json:"place"`
	Date 			time.Time						`json:"date"`
	Count 			int								`json:"count"`
	Fundraising 	int								`json:"fundraising"`
	UserId 			int								`json:"user_id"`

	AdditionalData 	[]AdditionalDataRowTypeResponse `json:"additional_data"`
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

	event, additionalDatas, err := handler.Event.GetEvent(eventId)
	if err != nil {
		responses.InternalServerError(c)
		handler.Logger.Debug(err.Error())
		return
	}

	var additionalDataResponse []AdditionalDataRowTypeResponse;
	for _, additionalData := range additionalDatas {
		additionalDataResponse = append(additionalDataResponse, 
			AdditionalDataRowTypeResponse{
				Id: additionalData.Id,
				AdditionalDataRowType: AdditionalDataRowType{
					Key: additionalData.Key,
					Value: additionalData.Value,
				},
				EventId: additionalData.EventId,
			},
		)
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

		AdditionalData: additionalDataResponse,
	})
}