package app

import (
	"net/http"

	"github.com/schattenbrot/go-simple-upload-server/packages/responder"
)

// StatusResponse represents the server's status.
// swagger:model
type StatusResponse struct {
	// Message represents the server status.
	Message string `json:"message"`
}

// @Summary Get server status
// @Description Retrieves the running server's status
// @Tags app
// @Produce json
// @Success 200 {object} StatusResponse
// @Router /app/status [get]
// @Router /app [get]
func status(w http.ResponseWriter, r *http.Request) {
	responder.Send(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "running",
	})
}

// @Summary Ping the server
// @Description Retrieves a success message
// @Tags app
// @Produce json
// @Success 200 {object} StatusResponse
// @Router /app/ping [get]
func ping(w http.ResponseWriter, r *http.Request) {
	responder.Send(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "successful",
	})
}
