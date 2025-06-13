package error_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(message string) (int, gin.H) {
	return http.StatusBadRequest, gin.H{"message": message}
}

func ServerError(message string) (int, gin.H) {
	return http.StatusInternalServerError, gin.H{"message": message}
}
