package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/tushar0305/event-management/db"
	"github.com/tushar0305/event-management/models"
)

func main(){
	db.InitDb()
	server := gin.Default()

	server.GET("/events", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080") //localhost
}

func getEvent(context *gin.Context){
	events, err := models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again Later!"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse"})
		return
	}

	event.Id = 1
	event.UserId = 1

	err = event.Save()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events. Try again Later!"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "events": event})
}