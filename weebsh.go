package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// WeebshHandler Handle all incoming weeb.sh API requests
func WeebshHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if !CheckIfAuthorized(w, r) {
		return
	}

	vars := mux.Vars(r)
	test := `
		{
			"resultsd": "` + vars["type"] + `"
		}
	`

	w.Write([]byte(test))
}
