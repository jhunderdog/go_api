package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jhunderdog/go_api/middlewares"
)

func RegisterRoutes(server *gin.Engine){
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)	
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
	// server.POST("/events", middlewares.Authenticate, createEvent)	
	// server.PUT("/events/:id", updateEvent)
	// server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}