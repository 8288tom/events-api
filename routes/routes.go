package routes

import (
	"example.com/events-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("events/:id/register", registerEvent)
	authenticated.DELETE("events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
