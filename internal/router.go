package router

import (
	"github.com/dmitriitalent/strittenApi/internal/handlers"
	authHandler "github.com/dmitriitalent/strittenApi/internal/handlers/auth"
	eventHandler "github.com/dmitriitalent/strittenApi/internal/handlers/event"
	userHandler "github.com/dmitriitalent/strittenApi/internal/handlers/user"
	"github.com/gin-gonic/gin"
)

type Router struct {
	authHandler.Auth
	userHandler.User
	eventHandler.Event
}

func NewRouter(handlers handlers.Handlers) *Router {
	return &Router{
		Auth: handlers.Auth,
		User: handlers.User,
		Event: handlers.Event,
	}
}

func (r *Router) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", r.Auth.Login)
			auth.POST("/registration", r.Auth.Registration)
			auth.POST("/refresh", r.Auth.Refresh)
			auth.POST("/logout", r.Auth.Logout)
		}

		user := api.Group("/user")
		{
			user.GET("/", r.User.GetUser)
			user.POST("/update", r.User.UpdateUser)
		}

		event := api.Group("/event")
		{
			event.GET("/", r.Event.GetEvent)
			event.POST("/create", r.Event.CreateEvent)
		}
	}

	return router
}

