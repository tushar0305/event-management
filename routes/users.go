package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/tushar0305/event-management/models"
)

func SignUp(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	_, err = user.Save()

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not save User!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user created sucessfully"})

}