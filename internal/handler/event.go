package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dmitriitalent/strittenApi/internal/dto"
	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createEvent(c *gin.Context) {
	accessTokenCookie, err := c.Request.Cookie("rt")
	if err != nil {
		errorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	token, err := h.services.Jwt.ValidateToken(accessTokenCookie.Value, []byte(saltRefreshToken))
	if err != nil {
		errorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	claims, err := h.services.Jwt.GetClaimsFromToken(token)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var userIdClaim string = "user_id"
	userId, ok := claims[userIdClaim].(int)

	if !ok {
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("%s claim not found at access token", userIdClaim).Error())
		return
	}

	var createEventData dto.CreateEvent;
	if err := c.BindJSON(&createEventData); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	Event := entity.Event{
		Name: createEventData.Name,
		Description: createEventData.Description,
		Place: createEventData.Place,
		Date: time.Now(),
		Count: createEventData.Count,
		Fundraising: createEventData.Fundraising,
		UserId: userId,
	}	
	h.services.CreateEvent(Event)
	
	c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Cool",
	})
}

func (h *Handler) getEvent(c *gin.Context) {


	c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Cool",
	})
}