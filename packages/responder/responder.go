package responder

import (
	"encoding/json"
	"net/http"
)

// Send is the helper function for sending back an HTTP response.
func Send(w http.ResponseWriter, status int, data interface{}) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

// Send is the helper function for sending back an HTTP response.
func SendFile(w http.ResponseWriter, file []byte) {
	fileHeader := file[:512]
	fileContentType := http.DetectContentType(fileHeader)

	w.Header().Set("Content-Type", fileContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}
