package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){
	server.GET("/events", GetEvents)
	server.POST("/events", CreateEvent)
	server.GET("/events/:id", GetEvent)
}