package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"mini-api/internal/scores"

	"github.com/go-chi/chi/v5"
)

func (app *application) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{ "api": "ok" }`))
}

func (app *application) Scores(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uid, ierr := strconv.Atoi(id)
	best, berr := strconv.Atoi(r.FormValue("b"))
	page, perr := strconv.Atoi(r.FormValue("p"))
	mode := r.FormValue("m")

	badParams := (ierr != nil || berr != nil || perr != nil)
	badArgs := (uid < -1 || best < 0 || best > 1 || page < -1)

	if badParams || badArgs || !scores.ValidGamemode(mode) {
		app.BadRequest(w)
		return
	}

	res, err := app.DB.GetScores(best != 0, uid, page, mode)
	scores, jerr := json.Marshal(res)

	if err != nil || jerr != nil {
		app.err.Println(err, jerr)
		app.InternalError(w)

		return
	}

	if len(res) <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(scores)
}

func (app *application) Score(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	sid, err := strconv.Atoi(id)

	if err != nil {
		app.BadRequest(w)
		return
	}

	res, err := app.DB.ScoreInfo(sid)
	scores, jerr := json.Marshal(res)

	if (err != nil && err != sql.ErrNoRows) || jerr != nil {
		app.err.Println(err, jerr)
		app.InternalError(w)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(scores)
}

func (app *application) Stats(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	mode := r.FormValue("m")
	uid, serr := strconv.Atoi(id)

	if serr != nil || !scores.ValidGamemode(mode) {
		app.BadRequest(w)
		return
	}

	res, err := app.DB.UserStats(uid, mode)
	stats, jerr := json.Marshal(res)

	if (err != nil && err != sql.ErrNoRows) || jerr != nil {
		app.err.Println(err, jerr)
		app.InternalError(w)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(stats)
}

func (app *application) Info(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uid, serr := strconv.Atoi(id)

	if serr != nil {
		app.BadRequest(w)
		return
	}

	res, err := app.DB.UserInfo(uid)
	info, jerr := json.Marshal(res)

	if (err != nil && err != sql.ErrNoRows) || jerr != nil {
		app.err.Println(err, jerr)
		app.InternalError(w)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(info)
}
