package eventHandler

import (
	"time"

	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/dmitriitalent/strittenApi/internal/responses"
	jwtService "github.com/dmitriitalent/strittenApi/internal/services/jwt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type CreateEventRequest struct {
	Name			string 					`json:"name"`
	Description		string 					`json:"description"`
	Place			string 					`json:"place"`
	Date			time.Time 				`json:"date"`
	Count			int 					`json:"count"`
	Fundraising		int						`json:"fundraising"`

	AdditionalData 	[]AdditionalDataRowType `json:"additional_data"`
}

type CreateEventResponse struct {
	Id          int			`json:"id"`
	Name        string		`json:"name"`
	Description string		`json:"description"`
	Place       string		`json:"place"`
	Date        time.Time	`json:"date"`
	Count       int			`json:"count"`
	Fundraising int			`json:"fundraising"`
	UserId      int			`json:"user_id"`

	AdditionalData 	[]AdditionalDataRowTypeResponse `json:"additional_data"`
}

func (handler *EventHandler) CreateEvent(c *gin.Context) {
	var dto CreateEventRequest;
	
	if err := c.BindJSON(&dto); err != nil {
		responses.BadRequest(c, "Some fields not found")
		handler.Logger.Debug(err.Error())
		return
	}

	refreshTokenCookie, err := c.Request.Cookie("rt")
	if err != nil {
		responses.Unauthorized(c, err.Error())
		handler.Logger.Debug(err.Error())
		return
	}

	token, err := handler.Jwt.Validate(refreshTokenCookie.Value, []byte(viper.GetString("crypto.refreshTokenSalt")))
	if err != nil {
		responses.Unauthorized(c, err.Error())
		handler.Logger.Debug(err.Error())
		return
	}

	claims, err := handler.Jwt.GetClaims(token)
	if err != nil {
		responses.InternalServerError(c)
		handler.Logger.Debug(err.Error())
		return
	}
	
	userId := int(claims[jwtService.UserIdClaim].(float64))

	additionalData := entity.AdditionalDatas{}
	for _, row := range dto.AdditionalData {
		additionalData = append(additionalData, entity.AdditionalData{
			Key:   row.Key,
			Value: row.Value,
		})
	}

	event := entity.Event{
		Name: 			dto.Name,
		Description: 	dto.Description,
		Place: 			dto.Place,
		Date: 			dto.Date,
		Count: 			dto.Count,
		Fundraising: 	dto.Fundraising,
		UserId:			userId,
	}

	createdEvent, createdAdditionalDatas, err := handler.Event.CreateEvent(event, additionalData);
	if err != nil {
		responses.InternalServerError(c)
		handler.Logger.Debug("CreateEventHandler:%s", err.Error())
		return 
	}

	var additionalDataResponse []AdditionalDataRowTypeResponse
	for _, createdAdditionalData := range createdAdditionalDatas {
		additionalDataResponse = append(additionalDataResponse, 
			AdditionalDataRowTypeResponse{
				Id: createdAdditionalData.Id,
				AdditionalDataRowType: AdditionalDataRowType{
					Key:   createdAdditionalData.Key,
					Value: createdAdditionalData.Value,
				},
				EventId: createdAdditionalData.EventId,
			},
		)
	}

	responses.Ok(c, CreateEventResponse{
		Id: 			createdEvent.Id,
		Name: 			createdEvent.Name,
		Description: 	createdEvent.Description,
		Place: 			createdEvent.Place,
		Date: 			createdEvent.Date,
		Count: 			createdEvent.Count,
		Fundraising: 	createdEvent.Fundraising,
		UserId: 		createdEvent.UserId,

		AdditionalData: additionalDataResponse,
	})
}