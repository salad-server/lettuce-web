package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) BadRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	w.Write([]byte(`{
		"code": 400,
		"error": "Bad request!"
	}`))
}

func (app *application) InternalError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	w.Write([]byte(`{
		"code": 500,
		"error": "Internal server error!"
	}`))
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
