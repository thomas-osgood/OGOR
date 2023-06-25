// package that provides general functions, types, structs, etc.
// that can be used when creating an API or HTTP server.
//
// inspiration for adding/based on: https://www.youtube.com/watch?v=pwZuNmAzaH8
package apis

import (
	"encoding/json"
	"net/http"
)

// function designed to more elegantly handle HTTP routing functions.
// this is, essentially, a middleware controller that extends the
// HandleFunc function. this takes an APIFunc as an argument and
// processes it.
func MakeHTTPHandleFunc(fnc APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error = fnc(w, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/text")
			w.Write([]byte(err.Error()))
			return
		}
	}
}

// function designed to handle writing a JSON payload to
// an HTTP response.
//
// if there is an issue encountered while encoding the
// JSON payload, the error will be returned, otherwise
// nil will be returned.
func WriteJSON(w http.ResponseWriter, status int, v any) (err error) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
