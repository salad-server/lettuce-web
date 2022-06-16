package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (app *application) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{ "api": "ok" }`))
}

func (app *application) Scores(w http.ResponseWriter, r *http.Request) {
	uid, ierr := strconv.Atoi(r.FormValue("u"))
	mode, merr := strconv.Atoi(r.FormValue("m"))
	best, berr := strconv.Atoi(r.FormValue("b"))
	page, perr := strconv.Atoi(r.FormValue("p"))

	badParams := (ierr != nil || merr != nil || berr != nil || perr != nil)
	badArgs := (uid < -1 || mode < -1 || best < 0 || best > 1 || page < -1)

	if badParams || badArgs {
		app.BadRequest(w)
		return
	}

	res, err := app.DB.GetScores(best != 0, uid, mode, page)
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
