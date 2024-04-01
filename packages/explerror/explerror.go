package explerror

import (
	"log"
	"net/http"
)

var logger *log.Logger
var send func(w http.ResponseWriter, logger *log.Logger, status int, data interface{}) error

type jsonError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func Setup(log *log.Logger, sendF func(w http.ResponseWriter, logger *log.Logger, status int, data interface{}) error) {
	logger = log
	send = sendF
}

func sendError(w http.ResponseWriter, statusCode int, err error) {
	logger.Printf("%d: %s", statusCode, err.Error())

	theError := &jsonError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}

	// sendError(w, statusCode, theError)
	send(w, logger, statusCode, theError)
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
