package explerror

import (
	"log"
	"net/http"
)

var logger *log.Logger
var send func(w http.ResponseWriter, status int, data interface{}) error

// ErrorResponse represents the structure for the error response.
// swagger:model
type ErrorResponse struct {
	// StatusCode represents the HTTP status code of the error.
	// example: 400
	StatusCode int `json:"statusCode"`

	// Message represents the error message.
	// example: Bad request
	Message string `json:"message"`
}

func Setup(loggerRef *log.Logger, sendF func(w http.ResponseWriter, status int, data interface{}) error) {
	logger = loggerRef
	send = sendF
}

func sendError(w http.ResponseWriter, statusCode int, err error) {
	logger.Printf("%d: %s", statusCode, err.Error())

	theError := &ErrorResponse{
		StatusCode: statusCode,
		Message:    err.Error(),
	}

	send(w, statusCode, theError)
}

func BadRequest(w http.ResponseWriter, err error) {
	sendError(w, http.StatusBadRequest, err)
}

func Forbidden(w http.ResponseWriter, err error) {
	sendError(w, http.StatusForbidden, err)
}

func InternalServerError(w http.ResponseWriter, err error) {
	sendError(w, http.StatusInternalServerError, err)
}

func NotFound(w http.ResponseWriter, err error) {
	sendError(w, http.StatusNotFound, err)
}

func Unauthorized(w http.ResponseWriter, err error) {
	sendError(w, http.StatusUnauthorized, err)
}
