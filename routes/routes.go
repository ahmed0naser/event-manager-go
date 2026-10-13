package routes

import (
	"github.com/ahmed0naser/event-manager-go/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){
server.GET("/events",getEvents)
server.GET("/events/:id",getEvent)
authenticated:=server.Group("/")
authenticated.Use(middlewares.Authenticate)
authenticated.POST("/events",createEvent)
authenticated.PUT("/events/:id",updateEvent)
authenticated.DELETE("/events/:id",deleteEvent)
authenticated.POST("events/:id/register",registerForEvent)
authenticated.DELETE("events/:id/register",cancelRegistration)
server.POST("/signup",Signup)
server.POST("/login",Login)
}