package routes

import (
	"fmt"
	"net/http"
	"strconv"

	error_handler "example.com/rest-api/errors"
	"example.com/rest-api/models"
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
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Print("Could not parse event id:", err)
		context.JSON(error_handler.BadRequest("Could not parse JSON data."))
		return
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
	err := context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Print("could not parse JSON:", err)
		context.JSON(error_handler.BadRequest("Could not parse JSON data."))
		return
	}

	event.UserID = 1

	err = event.Save()
	if err != nil {
		fmt.Print("could not save to database", err)
		context.JSON(error_handler.ServerError("Could not save to Database."))
		return

	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})

}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Print("Could not parse event id:", err)
		context.JSON(error_handler.BadRequest("Could not parse JSON data."))
		return
	}
	_, err = models.GetEventByID(eventId)

	if err != nil {
		fmt.Print("Could not fetch event", err)
		context.JSON(error_handler.ServerError("Could not fetch event."))
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		fmt.Print("Could not parse request data", err)
		context.JSON(error_handler.BadRequest("Could not parse request data."))
		return
	}
	updatedEvent.ID = eventId
	updatedEvent.Update()

	if err != nil {
		fmt.Print("Could not update event", err)
		context.JSON(error_handler.ServerError("Could not update event."))
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}
