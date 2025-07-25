package router

import (
	"github.com/dmitriitalent/strittenApi/internal/handlers"
	authHandler "github.com/dmitriitalent/strittenApi/internal/handlers/auth"
	userHandler "github.com/dmitriitalent/strittenApi/internal/handlers/user"
	"github.com/gin-gonic/gin"
)

type Router struct {
	authHandler.Auth
	userHandler.User
}

func NewRouter(handlers handlers.Handlers) *Router {
	return &Router{
		Auth: handlers.Auth,
		User: handlers.User,
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
		}
	}

	return router
}

