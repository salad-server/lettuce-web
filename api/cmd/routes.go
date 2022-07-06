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
	mux.Get("/leaderboard", app.Leaderboard)

	// auth
	mux.Post("/login", app.Login)
	// mux.Post("/register", app.Register)

	mux.Route("/@me", func(r chi.Router) {
		r.Use(app.IsAuthed)
		r.Get("/", app.WhoAmI)

		r.Get("/profile", app.GetProfile)
		r.Post("/profile", app.UpdateProfile)

		r.Post("/pfp", app.ProfilePicture)
		r.Delete("/pfp", app.ProfilePictureRemove)
	})

	// docs
	mux.Get("/docs", app.DocsListing)
	mux.Get("/docs/", app.DocsListing)
	mux.Handle("/docs/*", app.Docs())

	// users
	mux.Get("/users/{id}", app.Info)
	mux.Get("/users/{id}/scores", app.Scores)
	mux.Get("/users/{id}/stats", app.Stats)

	// scores
	mux.Get("/score", app.Records)
	mux.Get("/score/{id}", app.Score)

	// beatmap
	mux.Get("/beatmap/{id}", app.Beatmap)
	mux.Get("/beatmap/{id}/sets", app.BeatmapSets)
	mux.Get("/beatmap/{id}/leaderboard", app.BeatmapLeaderboard)

	return mux
}
