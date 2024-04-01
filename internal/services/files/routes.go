package files

import "github.com/go-chi/chi/v5"

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", uploadFile)
	r.Get("/", getFiles)
	r.Get("/{filename}", getFile)

	return r
}
