package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(app.cors))

	mux.Get("/", app.Index)

	// users
	mux.Get("/users/{id}", app.Info)
	mux.Get("/users/{id}/scores", app.Scores)
	mux.Get("/users/{id}/stats", app.Stats)
	
	// scores
	mux.Get("/score/{id}", app.Score)

	// beatmap
	mux.Get("/beatmap/{id}", app.Beatmap)
	mux.Get("/beatmap/{id}/sets", app.BeatmapSets)
	
	return mux
}
