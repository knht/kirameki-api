package main

import "net/http"

// CheckIfAuthorized Check if a request is authorized.
func CheckIfAuthorized(w http.ResponseWriter, r *http.Request) bool {
	if r.Header.Get("Authorization") == "" {
		w.Write([]byte(`{ "error": "Unauthorized. Please make sure your authorization token is set correctly." }`))
		return false
	} else if r.Header.Get("Authorization") != "" && !Contains(authtokens, r.Header.Get("Authorization")) {
		w.Write([]byte(`{ "error": "Unauthorized. Please make sure your authorization token is set correctly." }`))
		return false
	}

	return true
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
