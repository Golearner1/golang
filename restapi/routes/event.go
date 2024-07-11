package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/first-app/models"
	"github.com/gin-gonic/gin"
)

func getevents(context *gin.Context) {
	event, err := models.Getallevents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events "})
		return
	}
	context.JSON(http.StatusOK, event)
}
func getevent(context *gin.Context) {
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
	context.JSON(http.StatusOK, event)
}
func createevents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}
	authid := context.GetInt64("authid")
	event.Authid = int64(authid)
	err = event.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create events "})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})

}
func updateevent(context *gin.Context) {
	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not convert id "})
		return
	}
	authid := context.GetInt64("authid")
	event, err := models.GetEventByID(eventid)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	if event.Authid != authid {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update event"})
		return
	}
	var updatedevent models.Event
	err = context.ShouldBindJSON(&updatedevent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}
	updatedevent.ID = eventid
	err = updatedevent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"mrssage": "event has updated successfully"})
}
func deleteevent(context *gin.Context) {
	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not convert id "})
		return
	}
	authid := context.GetInt64("authid")
	events, err := models.GetEventByID(eventid)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	if events.Authid != authid {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete event"})
		return
	}
	err = events.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event has deleted successfully"})
}
