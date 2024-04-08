package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/schattenbrot/go-simple-upload-server/internal/config"
	"github.com/schattenbrot/go-simple-upload-server/internal/services/app"
	"github.com/schattenbrot/go-simple-upload-server/internal/services/files"
	"github.com/schattenbrot/go-simple-upload-server/packages/explerror"
	"github.com/schattenbrot/go-simple-upload-server/packages/responder"
)

func main() {
	config.NewConfig()
	explerror.Setup(log.Default(), responder.Send)

	// Check if the files-directory exists
	directory := "./data/files"
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		// If the directory doesn't exist, create it
		err := os.MkdirAll(directory, 0755)
		if err != nil {
			log.Fatal("Error creating directory:", err)
			return
		}
		fmt.Println("File Directory created successfully")
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:     config.Cors.AllowedOrigins,
		OptionsPassthrough: true,
	}))

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/", app.Routes())
		r.Mount("/files", files.Routes())
	})

	log.Println("The API runs on port:", config.Port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r); err != nil {
		log.Println(err.Error())
	}
}
