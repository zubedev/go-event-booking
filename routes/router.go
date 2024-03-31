package routes

import (
	"github.com/gin-gonic/gin"
	"go-event-booking/middlewares"
)

func RegisterRouter(server *gin.Engine) {
	// users routes
	server.POST("/users/signup", signup)
	server.POST("/users/login", login)

	// events routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// authenticated routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PATCH("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

}
