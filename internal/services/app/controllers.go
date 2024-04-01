package app

import (
	"net/http"

	"github.com/schattenbrot/go-simple-upload-server/packages/responder"
)

func status(w http.ResponseWriter, r *http.Request) {
	responder.Send(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "running",
	})
}
