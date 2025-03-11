package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/api/measurement", app.GetAllMeasurements)
	mux.Post("/api/measurement/create", app.CreateMeasurement)

	return mux
}
