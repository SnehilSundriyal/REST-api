package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // we are setting up a handler for an incoming GET request
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
}