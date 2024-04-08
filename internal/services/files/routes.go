package files

import (
	"github.com/go-chi/chi/v5"
)

func Routes() chi.Router {
	r := chi.NewRouter()

	r.With(hasReadWriteAccess).Post("/", uploadFile)
	r.With(hasReadAccess).Get("/", getFiles)
	r.With(hasReadAccess).Get("/{filename}", getFile)

	return r
}
