package handlers

import (
	"errors"
	"net/http"
	"server/app/services"

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

// SendError handles different types of errors and sends appropriate HTTP responses.
//
// This function checks the type of the error and calls the corresponding error
// handling function with the appropriate status code and message.
//
// Parameters:
//   - err: The error to be handled.
//   - c: The Gin context for the current request.
func SendError(err error, c *gin.Context) {
	var entityNotFoundError services.EntityNotFoundError
	var createEntityFailureError services.CreateEntityFailureError
	var permissionDeniedError services.PermissionDeniedError
	var cannotPerformActionError services.CannotPerformActionError

	switch {
	case errors.As(err, &entityNotFoundError):
		HandleNotFound(c, err.Error())
		return

	case errors.As(err, &permissionDeniedError) || errors.As(err, &cannotPerformActionError):
		HandleUnauthorized(c, err.Error())
		return

	case errors.As(err, &createEntityFailureError):
		HandleError(c, http.StatusInternalServerError, err.Error())

	default:
		HandleError(c, http.StatusInternalServerError, "an error occurred")
	}
}
