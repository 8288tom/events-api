package helpers

import (
	"fmt"
	"strconv"

	error_handler "example.com/events-api/errors"
	"github.com/gin-gonic/gin"
)

func GetEventId(context *gin.Context) (int64, error) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Print("Could not parse event id:", err)
		context.JSON(error_handler.BadRequest("Could not parse JSON data."))
		return 0, err
	}
	return eventId, nil
}
func GetJSONData(context *gin.Context, data any) error {
	err := context.ShouldBindJSON(data)
	if err != nil {
		fmt.Print("could not parse JSON:", err)
		context.JSON(error_handler.BadRequest("Could not parse JSON data."))
		return err
	}
	return nil
}

func GetToken(context *gin.Context) string {
	token := context.Request.Header.Get("Authorization")
	return token
}
