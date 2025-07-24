package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dmitriitalent/strittenApi/internal/dto"
	"github.com/dmitriitalent/strittenApi/internal/entity"
	"github.com/dmitriitalent/strittenApi/internal/service"
	"github.com/gin-gonic/gin"
)

const (
	saltAccessToken  = "1"
	saltRefreshToken = "2"
)

func (h *Handler) login(c *gin.Context) {
	var loginData dto.Login

	if err := c.BindJSON(&loginData); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Validation.IsPasswordValid(loginData.Password); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var user = entity.User{
		Login:    loginData.Login,
		Password: loginData.Password,
	}

	id, err := h.services.User.FindUser(user)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, err := h.services.Jwt.GenerateAccessToken(service.TokenClaims{
		UserId: id,
	}, "1")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := h.services.Jwt.GenerateRefreshToken(service.TokenClaims{
		UserId: id,
	}, "2")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.SetCookie(
		"at",
		accessToken,
		int(time.Now().Add(10*time.Minute).Unix()),
		"/",
		"",
		true,
		false,
	)

	c.SetCookie(
		"rt",
		refreshToken,
		int(time.Now().Add(30*time.Hour*24).Unix()),
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, map[string]interface{}{
		"at": accessToken,
	})
}

func (h *Handler) registration(c *gin.Context) {
	var registrationData dto.Registration

	if err := c.BindJSON(&registrationData); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Validation.IsEmailValid(registrationData.Email); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if registrationData.Password != registrationData.ConfirmPassword {
		errorResponse(c, http.StatusBadRequest, "Passwords do not match")
		return
	}

	if err := h.services.Validation.IsPasswordValid(registrationData.Password); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var user = entity.User{
		Login:    registrationData.Login,
		Password: registrationData.Password,
		Name: 	  registrationData.Name,
		Surname:  registrationData.Surname,
		Email:    registrationData.Email,
	}

	id, err := h.services.User.CreateUser(user)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	accessToken, err := h.services.Jwt.GenerateAccessToken(service.TokenClaims{
		UserId: id,
	}, "1")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := h.services.Jwt.GenerateRefreshToken(service.TokenClaims{
		UserId: id,
	}, "2")
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.SetCookie(
		"at",
		accessToken,
		int(time.Now().Add(10*time.Minute).Unix()),
		"/",
		"",
		true,
		false,
	)

	c.SetCookie(
		"rt",
		refreshToken,
		int(time.Now().Add(30*time.Hour*24).Unix()),
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, map[string]interface{}{
		"at": accessToken,
	})
}

func (h *Handler) refresh(c *gin.Context) {
	refreshTokenCookie, err := c.Request.Cookie("rt")
	if err != nil {
		errorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	token, err := h.services.Jwt.ValidateToken(refreshTokenCookie.Value, []byte(saltRefreshToken))
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
	userId, ok := claims[userIdClaim].(float64)

	if !ok {
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("%s claim not found at refresh token", userIdClaim).Error())
		return
	}

	accessToken, err := h.services.Jwt.GenerateAccessToken(service.TokenClaims{
		UserId: int(userId),
	}, saltAccessToken)

	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := h.services.Jwt.GenerateRefreshToken(service.TokenClaims{
		UserId: int(userId),
	}, saltRefreshToken)

	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.SetCookie(
		"at",
		accessToken,
		int(time.Now().Add(10*time.Minute).Unix()),
		"/",
		"",
		true,
		false,
	)

	c.SetCookie(
		"rt",
		refreshToken,
		int(time.Now().Add(30*time.Hour*24).Unix()),
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, map[string]interface{}{
		"at": accessToken,
	})
}

func (h *Handler) logout(c *gin.Context) {
	refreshTokenCookie, err := c.Request.Cookie("rt")
	if err != nil {
		errorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	_, err = h.services.Jwt.ValidateToken(refreshTokenCookie.Value, []byte(saltRefreshToken))
	if err != nil {
		errorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	c.SetCookie("at", "", 0, "", "", false, false)
	c.SetCookie("rt", "", 0, "", "", false, false)

	h.services.Jwt.RemoveRefreshToken(refreshTokenCookie.Value)
}

type getUserParams struct {
	UserId int `form:"userid"`
}

func (h *Handler) getUser(c *gin.Context) {
	var getUserParams getUserParams

	if err := c.ShouldBindQuery(&getUserParams); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.FindUserById(getUserParams.UserId)

	if err != nil {
		errorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
