package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tushar0305/event-management/models"
)

func RegisterForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id!"})
        return
    }

	_, err = models.GetEventById(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch the event!"})
		return
	}

	// Check if the user is already registered for the event
	isRegistered, err := models.IsUserRegisteredForEvent(userId, eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not check registration status!"})
		return
	}
	if isRegistered {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User already registered for the event!"})
		return
	}

	// Register the user for the event
	err = models.RegisterUserForEvent(userId, eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for the event!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User registered for the event successfully!"})
}

func CancelRegistration(context *gin.Context) {
    userId := context.GetInt64("userId")
    eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id!"})
        return
    }

    err = models.CancelUserRegistration(userId, eventId)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration!"})
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "Registration canceled successfully!"})
}