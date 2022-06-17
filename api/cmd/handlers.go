package main

import (
	"net/http"
	"strconv"

	"mini-api/internal/scores"

	"github.com/go-chi/chi/v5"
)

func (app *application) Index(w http.ResponseWriter, r *http.Request) {
	if app.conf.env == "development" {
		app.JSON(w, devResponse{
			Mode:   app.conf.env,
			Front:  app.conf.cors[0],
			Uptime: app.getUptime(),
		})

		return
	}

	http.Redirect(w, r, app.conf.cors[0], http.StatusPermanentRedirect)
}

func (app *application) Scores(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uid, ierr := strconv.Atoi(id)
	best, berr := strconv.Atoi(r.FormValue("b"))
	page, perr := strconv.Atoi(r.FormValue("p"))
	mode := r.FormValue("m")

	badParams := (ierr != nil || berr != nil || perr != nil)
	badArgs := (uid <= -1 || best < 0 || best > 1 || page <= -1)

	if badParams || badArgs || !scores.ValidGamemode(mode) {
		app.BadRequest(w)
		return
	}

	res, err := app.DB.GetScores(best != 0, uid, page, mode)

	if err != nil {
		app.err.Println(err)
		app.InternalError(w)

		return
	}

	app.JSON(w, res)
}

func (app *application) Score(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	sid, err := strconv.Atoi(id)

	if err != nil {
		app.BadRequest(w)
		return
	}

	res, err := app.DB.ScoreInfo(sid)

	if err != nil {
		app.err.Println(err)
		app.InternalError(w)

		return
	}

	app.JSON(w, res)
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

	if err != nil {
		app.err.Println(err)
		app.InternalError(w)

		return
	}

	app.JSON(w, res)
}

func (app *application) Info(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uid, serr := strconv.Atoi(id)

	if serr != nil {
		app.BadRequest(w)
		return
	}

	res, err := app.DB.UserInfo(uid)

	if err != nil {
		app.err.Println(err)
		app.InternalError(w)

		return
	}

	app.JSON(w, res)
}

func (app *application) Beatmap(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bid, serr := strconv.Atoi(id)

	if serr != nil {
		app.BadRequest(w)
		return
	}

	res, err := app.DB.BeatmapInfo(bid)

	if err != nil {
		app.err.Println(err)
		app.InternalError(w)

		return
	}

	app.JSON(w, res)
}

func (app *application) BeatmapSets(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bid, serr := strconv.Atoi(id)

	if serr != nil {
		app.BadRequest(w)
		return
	}

	res, err := app.DB.BeatmapSets(bid)

	if err != nil {
		app.err.Println(err)
		app.InternalError(w)

		return
	}

	app.JSON(w, res)
}

func (app *application) Records(w http.ResponseWriter, r *http.Request) {
	res, err := app.DB.Records()

	if err != nil {
		app.err.Println(err)
		app.InternalError(w)
		return
	}

	app.JSON(w, res)
}

func (app *application) BeatmapLeaderboard(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bid, berr := strconv.Atoi(id)
	page, perr := strconv.Atoi(r.FormValue("p"))
	mode := r.FormValue("m")

	badParams := berr != nil || perr != nil || !scores.ValidGamemode(mode)
	badArgs := (bid <= -1 || page <= -1)

	if badParams || badArgs {
		app.BadRequest(w)
		return
	}

	res, err := app.DB.BeatmapLeaderboard(bid, page, mode)

	if err != nil {
		app.err.Println(err)
		app.InternalError(w)

		return
	}

	app.JSON(w, res)
}

func (app *application) Leaderboard(w http.ResponseWriter, r *http.Request) {
	page, perr := strconv.Atoi(r.FormValue("p"))
	sort, serr := strconv.Atoi(r.FormValue("s"))
	mode := r.FormValue("m")

	if perr != nil || serr != nil || page <= -1 || !scores.ValidGamemode(mode) {
		app.BadRequest(w)
		return
	}

	// 0 = pp
	// 1 = score
	res, err := app.DB.Leaderboard(sort != 1, mode, page)

	if err != nil {
		app.err.Println(err)
		app.InternalError(w)

		return
	}

	app.JSON(w, res)
}
