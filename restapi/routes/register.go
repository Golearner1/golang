package routes

import (
	"net/http"
	"strconv"

	"example.com/first-app/models"
	"github.com/gin-gonic/gin"
)

func Eventregistration(context *gin.Context) {
	authid := context.GetInt64("authid")
	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not convert id "})
		return
	}
	event, err := models.GetEventByID(eventid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(authid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for the event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user has registered for the event successfully"})
}
func Cancelregistration(context *gin.Context) {
	authid := context.GetInt64("authid")
	eventid, _ := strconv.ParseInt(context.Param("id"), 10, 64)
	var event models.Event
	event.ID = eventid
	err := event.Deleteregistration(authid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "user has canceled for the event successfully"})
}
