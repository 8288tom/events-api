package routes

import (
	"fmt"
	"net/http"

	error_handler "example.com/events-api/errors"
	"example.com/events-api/helpers"
	"example.com/events-api/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Print("could not parse JSON:", err)
		context.JSON(error_handler.BadRequest("Could not parse JSON data."))
		return
	}

	err = user.Save()
	if err != nil {
		fmt.Print(err)
		context.JSON(error_handler.BadRequest("Could not save user."))
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}

func login(context *gin.Context) {
	var user models.User
	err := helpers.GetJSONData(context, &user)

	if err != nil {
		return //helper handles return
	}
	err = user.ValidateCredentials()

	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}

	token, err := helpers.GenerateToken(user.Email, user.ID)

	if err != nil {
		fmt.Print(err)
		context.JSON(error_handler.ServerError("Could not generate token"))
		return
	}

	context.SetCookie("token", token, 30, "/", "localhost", true, true)
	context.JSON(http.StatusOK, gin.H{"message": "Login successfull!"})

}
