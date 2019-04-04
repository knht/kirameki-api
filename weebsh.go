package main

import (
	"fmt"
	"net/http"

	weebgo "github.com/Daniele122898/weeb.go/src"
	"github.com/Daniele122898/weeb.go/src/data"
	"github.com/gorilla/mux"
)

// WeebAuth authenticate with the weeb.sh API once
func WeebAuth() error {
	err := weebgo.Authenticate(weebToken, data.WOLKE)
	if err == nil {
		return err
	}

	return err
}

// WeebshHandler Handle all incoming weeb.sh API requests
func WeebshHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if !CheckIfAuthorized(w, r) {
		return
	}

	vars := mux.Vars(r)
	ri, err := weebgo.GetRandomImage(vars["type"], nil, data.ANY, data.FALSE, false)

	if err != nil {
		fmt.Println(err)
		return
	}

	url := ri.Url

	result := `
		{
			"url": "` + url + `"
		}
	`

	w.Write([]byte(result))
}
