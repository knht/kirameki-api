package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
)

var authtokens = []string{}

// Config struct for auth tokens
type Config struct {
	Tokens []string `json:"tokens"`
}

// TestHandler easy test handler for mux
func TestHandler(w http.ResponseWriter, r *http.Request) {
	// Set application header to JSON
	w.Header().Add("Content-Type", "application/json")

	// Check for adequate authorization
	if r.Header.Get("Authorization") == "" {
		w.Write([]byte(`{ "error": "Unauthorized. Please make sure your authorization token is set correctly." }`))
		return
	} else if r.Header.Get("Authorization") != "" && !Contains(authtokens, r.Header.Get("Authorization")) {
		w.Write([]byte(`{ "error": "Unauthorized. Please make sure your authorization token is set correctly." }`))
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

// Contains simple function to check if an element exists within a slice
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}

	return false
}

func main() {
	// Read the config file
	configFile, err := os.Open("config.json")

	// Check if the config file was read successfully
	if err != nil {
		fmt.Println(err)
	}

	// Defer closing the file until parsed
	defer configFile.Close()

	// Read the config file
	byteValue, _ := ioutil.ReadAll(configFile)

	var tokens Config

	json.Unmarshal(byteValue, &tokens)

	// Append each auth token to the global auth token slice
	for i := 0; i < len(tokens.Tokens); i++ {
		authtokens = append(authtokens, tokens.Tokens[i])
	}

	// Instantiate new mux router
	r := mux.NewRouter()

	// Handle a simple test route
	r.HandleFunc("/{one}/{two}", TestHandler).Methods("GET")

	// Start listening for incoming request and logging if something is wrong
	log.Fatal(http.ListenAndServe(":8000", r))
}
