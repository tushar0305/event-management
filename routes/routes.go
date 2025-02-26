package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/tushar0305/event-management/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
    server.GET("/events", GetEvents)
    server.GET("/events/:id", GetEvent)

    authenticated := server.Group("/")
    authenticated.Use(middlewares.Authenticate)

    authenticated.POST("/events", CreateEvent)
    authenticated.PUT("/events/:id", UpdateEvent)
    authenticated.DELETE("/events/:id", DeleteEvent)
	authenticated.POST("/events/:id/register", RegisterForEvent)
	authenticated.DELETE("/events/:id/register", CancelRegistration)


    server.POST("/signup", SignUp)
    server.POST("/login", Login)
}