package routes

import (
	"fmt"
	"net/http"

	error_handler "example.com/events-api/errors"
	"example.com/events-api/helpers"
	"example.com/events-api/models"
	"github.com/gin-gonic/gin"
)

func registerEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := helpers.GetEventId(context)
	if err != nil {
		return // helper function handles response
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		fmt.Print(err)
		context.JSON(error_handler.BadRequest("Could not fetch event"))
		return
	}

	err = event.Register(userId)
	if err != nil {
		fmt.Print(err)
		context.JSON(error_handler.ServerError("Could not register user for event"))
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registed for event!"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := helpers.GetEventId(context)
	if err != nil {
		return // helper function handles response
	}
	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		fmt.Print(err)
		context.JSON(error_handler.ServerError("Could not cancel registration"))
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Cancelled!"})
}
