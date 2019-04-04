package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/gorilla/mux"
)

var authtokens = []string{}
var weebToken string
var port = "28785"

// Config struct for auth tokens
type Config struct {
	Tokens      []string `json:"tokens"`
	WeebshToken string   `json:"weebshToken"`
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

	// Set weeb.sh config token
	weebToken = tokens.WeebshToken

	// Instantiate new mux router
	r := mux.NewRouter()

	// Handle a simple test route
	r.HandleFunc("/test/{one}/{two}", TestHandler).Methods("GET")

	// Handle weeb.sh API
	r.HandleFunc("/weebsh/{type}", WeebshHandler).Methods("GET")

	// Boot up logo & information
	bootLogo := figure.NewFigure("kirAPI", "speed", true)
	bootLogo.Print()
	fmt.Println("\nNow running on port " + port)

	// Authenticate with weeb.sh
	WeebAuth()

	// Start listening for incoming request and logging if something is wrong
	log.Fatal(http.ListenAndServe(":"+port, r))
}
