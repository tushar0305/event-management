package routes

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/tushar0305/event-management/models"
)

func GetEvents(context *gin.Context){
	events, err := models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again Later!"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func CreateEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	userId := context.GetInt64(("userId"))
	event.UserId = userId

	err = event.Save()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event created successfully!"})
}

func GetEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id!"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event id!"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func UpdateEvent(context *gin.Context) {
    eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id!"})
        return
    }

    _, err = models.GetEventById(eventId)
    if err != nil {
        context.JSON(http.StatusNotFound, gin.H{"message": "Event not found!"})
        return
    }

    var updatedEvent models.Event
    err = context.ShouldBindJSON(&updatedEvent)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse requested data"})
        return
    }

    updatedEvent.Id = eventId

    err = updatedEvent.UpdateEventById()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event!"})
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id!"})
        return
    }

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found!"})
		return
	}

	err = event.DeleteEventById()
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not Delete the event!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})

}