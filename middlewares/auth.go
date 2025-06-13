package middlewares

import (
	"fmt"
	"net/http"

	"example.com/events-api/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := helpers.GetToken(context)

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := helpers.VerifyToken(token)
	if err != nil {
		fmt.Print(err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
