package routes

import (
	"fmt"
	"net/http"

	error_handler "example.com/events-api/errors"
	"example.com/events-api/helpers"
	"example.com/events-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Print("Failed getting all events", err)
		context.JSON(error_handler.ServerError("Failed getting all events"))
		return
	}
	context.JSON(200, events)
}

func getEvent(context *gin.Context) {
	eventId, err := helpers.GetEventId(context)
	if err != nil {
		return // helper function handles response
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		fmt.Printf("Failed getting event id:%v\n %v", eventId, err)
		context.JSON(error_handler.ServerError("Failed getting event"))
		return
	}
	context.JSON(200, event)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := helpers.GetJSONData(context, &event)
	if err != nil {
		return // helper function handles response
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		fmt.Print("could not save to database", err)
		context.JSON(error_handler.ServerError("Could not save to Database."))
		return

	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})

}

func updateEvent(context *gin.Context) {
	eventId, err := helpers.GetEventId(context)
	if err != nil {
		return // helper function handles response
	}

	event, err := models.GetEventByID(eventId)
	userId := context.GetInt64("userId")

	if err != nil {
		fmt.Print("Could not fetch event", err)
		context.JSON(error_handler.ServerError("Could not fetch event."))
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return
	}

	var updatedEvent models.Event

	err = helpers.GetJSONData(context, &updatedEvent)
	if err != nil {
		return // helper function handles response
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		fmt.Print("Could not update event", err)
		context.JSON(error_handler.ServerError("Could not update event."))
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := helpers.GetEventId(context)
	if err != nil {
		return // helper function handles response
	}

	event, err := models.GetEventByID(eventId)
	userId := context.GetInt64("userId")

	if err != nil {
		fmt.Print("Could not fetch event", err)
		context.JSON(error_handler.ServerError("Could not fetch event."))
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event"})
		return
	}
	err = event.Delete()

	if err != nil {
		context.JSON(error_handler.ServerError("Could not delete the event."))
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
