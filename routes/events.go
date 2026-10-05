package routes

import (
	"net/http"
	"strconv"

	"github.com/ahmed0naser/event-manager-go/db"
	"github.com/ahmed0naser/event-manager-go/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvants()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)

}
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse event id"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't fetch event"})
		return
	}
	context.JSON(http.StatusOK, event)

}
func createEvent(context *gin.Context) {
	db.InitDb()
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed to parse input"})
		return
	}
	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't create event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"status": "success", "event": event})

}
func updateEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse event id"})
		return
	}
	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't fetch event"})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed to parse input"})
		return
	}
	updatedEvent.ID=eventId
err=updatedEvent.Update()
if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't update event"})
		return
	}
		context.JSON(http.StatusOK, gin.H{"status": "event updated successfully"})

}
func deleteEvent(context *gin.Context){
eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse event id"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't fetch event"})
		return
	}
	err=event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't delete event"})
		return
	}
	
		context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
		
	

}