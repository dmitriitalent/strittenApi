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
	Name		string 		`json:"name"`
	Description	string 		`json:"description"`
	Place		string 		`json:"place"`
	Date		time.Time 	`json:"date"`
	Count		int 		`json:"count"`
	Fundraising	int			`json:"fundraising"`
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

	event := entity.Event{
		Name: 			dto.Name,
		Description: 	dto.Description,
		Place: 			dto.Place,
		Date: 			dto.Date,
		Count: 			dto.Count,
		Fundraising: 	dto.Fundraising,
		UserId:			userId,
	}

	createdEvent, err := handler.Event.CreateEvent(event);
	if err != nil {
		responses.InternalServerError(c)
		handler.Logger.Debug("CreateEventHandler:%s", err.Error())
		return 
	}

	responses.Ok(c, createdEvent)
}