package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var uptime time.Time

type errorRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (app *application) BadRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	app.JSON(w, errorRes{
		Code:    400,
		Message: "Bad request!",
	})
}

func (app *application) InternalError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	app.JSON(w, errorRes{
		Code:    500,
		Message: "Internal server error!",
	})
}

func (app *application) JSON(w http.ResponseWriter, jdata any) {
	j, err := json.Marshal(jdata)

	if err != nil {
		app.err.Println("Could not parse data!")
		app.InternalError(w)

		return
	}

	if string(j) == "null" {
		j = []byte("[]")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func (app *application) getUptime() string {
	return time.Since(uptime).String()
}

func init() {
	uptime = time.Now()
}
