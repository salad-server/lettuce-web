package main

import "net/http"

func (app *application) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{ "api": "ok" }`))
}
