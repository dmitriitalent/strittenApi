package handler

import (
	"github.com/dmitriitalent/strittenApi/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", h.login)
			auth.POST("/registration", h.registration)
			auth.POST("/refresh", h.refresh)
			auth.POST("/logout", h.logout)
		}

		user := api.Group("/user")
		{
			user.GET("", h.getUser)
		}

		event := api.Group("/event")
		{
			event.GET("", h.getEvent);
			event.POST("/create", h.createEvent);
		}
	}

	return router
}
