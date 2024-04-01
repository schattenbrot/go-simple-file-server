package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/schattenbrot/go-simple-upload-server/internal/config"
	"github.com/schattenbrot/go-simple-upload-server/internal/services/app"
	"github.com/schattenbrot/go-simple-upload-server/internal/services/files"
)

func main() {
	config.NewConfig()

	// Check if the files-directory exists
	directory := "./data/files"
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		// If the directory doesn't exist, create it
		err := os.MkdirAll(directory, 0755)
		if err != nil {
			log.Fatal("Error creating directory:", err)
			return
		}
		fmt.Println("Directory created successfully:", directory)
	}

	r := chi.NewRouter()

	r.Mount("/", app.Routes())
	r.Mount("/files", files.Routes())

	fmt.Println("Runs")

	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", config.Port), r); err != nil {
		fmt.Println(err.Error())
	}
}
