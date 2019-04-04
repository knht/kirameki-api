package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

// TestHandler easy test handler for mux
func TestHandler(w http.ResponseWriter, r *http.Request) {
	// Set application header to JSON
	w.Header().Add("Content-Type", "application/json")

	if !CheckIfAuthorized(w, r) {
		return
	}

	// Get parameters and execute Node.js worker
	vars := mux.Vars(r)
	out, err := exec.Command("node", "workers/test.js", vars["one"], vars["two"]).Output()

	// Check if the worker succeeded
	if err != nil {
		log.Fatal(err)
	}

	workerResult := `{
		"result": "` + string(out[:]) + `"
	}`

	// Send the result
	w.Write([]byte(workerResult))
}
