package routes

import "github.com/gin-gonic/gin"

func RegisterRouter(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PATCH("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
}
