// package that provides general functions, types, structs, etc.
// that can be used when creating an API or HTTP server.
//
// inspiration for adding/based on: https://www.youtube.com/watch?v=pwZuNmAzaH8
package apis

import (
	"encoding/json"
	"fmt"
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

// function designed to return an error JSON payload to the client.
// this can be passed to MakeHTTPHandleFunc and used if/when an
// issue occurs during the execution of an API endpoint.
func ReturnErrorJSON(w *http.ResponseWriter, status int, message string) (err error) {
	var payload ErrorStruct = ErrorStruct{}

	payload.ErrorCode = status
	payload.ErrorMessage = message

	return WriteJSON(w, status, &payload)
}

// function designed to handle writing a JSON payload to
// an HTTP response.
//
// if there is an issue encountered while encoding the
// JSON payload, the error will be returned, otherwise
// nil will be returned.
func WriteJSON(w *http.ResponseWriter, status int, v any) (err error) {
	var contentlength int

	(*w).Header().Set("Status-Code", fmt.Sprintf("%d", status))
	(*w).Header().Set("Content-Type", "application/json")

	contentlength = len(fmt.Sprintf("%v", v))

	(*w).Header().Set("Content-Length", fmt.Sprintf("%d", contentlength))

	return json.NewEncoder(*w).Encode(v)
}
