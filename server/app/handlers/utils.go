package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) uint {
	return c.GetUint("userId")
}

// HandleError sends a JSON response with the given status code and error message.
//
// Parameters:
//   - c: The Gin context for the current request.
//   - status: The HTTP status code to be sent in the response.
//   - message: The error message to be included in the response.
func HandleError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

// HandleBadRequest sends a JSON response with a 400 Bad Request status and the given error message.
//
// Parameters:
//   - c: The Gin context for the current request.
//   - message: The error message to be included in the response.
func HandleBadRequest(c *gin.Context, message string) {
	HandleError(c, http.StatusBadRequest, message)
}

// HandleNotFound sends a JSON response with a 404 Not Found status and the given error message.
//
// Parameters:
//   - c: The Gin context for the current request.
//   - message: The error message to be included in the response.
func HandleNotFound(c *gin.Context, message string) {
	HandleError(c, http.StatusNotFound, message)
}

// HandleUnauthorized sends a JSON response with a 401 Unauthorized status and the given error message.
//
// Parameters:
//   - c: The Gin context for the current request.
//   - message: The error message to be included in the response.
func HandleUnauthorized(c *gin.Context, message string) {
	HandleError(c, http.StatusUnauthorized, message)
}
