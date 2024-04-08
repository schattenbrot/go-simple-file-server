package app

import "github.com/go-chi/chi/v5"

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", status)
	r.Get("/status", status)
	r.Get("/ping", ping)

	return r
}
