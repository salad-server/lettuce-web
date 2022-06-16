package main

import "net/http"

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
