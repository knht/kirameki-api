package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var authtokens = []string{}

// Config struct for auth tokens
type Config struct {
	Tokens []string `json:"tokens"`
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
	r.HandleFunc("/test/{one}/{two}", TestHandler).Methods("GET")

	// Handle weeb.sh API
	r.HandleFunc("/weebsh/{type}", WeebshHandler).Methods("GET")

	// Start listening for incoming request and logging if something is wrong
	log.Fatal(http.ListenAndServe(":8000", r))
}
