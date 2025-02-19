package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tushar0305/event-management/models"
)

func main(){
	server := gin.Default()

	server.GET("/events", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080") //localhost
}

func getEvent(context *gin.Context){
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse"},)
		return
	}

	event.Id = 1
	event.UserId = 1

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "events": event})
}