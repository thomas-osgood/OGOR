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
			w.Header().Set("Status-Code", fmt.Sprintf("%d", http.StatusBadRequest))
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(err.Error()))
			return
		}
	}
}

// function designed to create and return a new MiddlwareController object
// to the user. this will return a pointer to the new MiddlwareController
// and an error. if the creation is successful, nil will be returned in
// place of an error.
func NewMiddlwareController(optsfuncs ...MiddlewareOptsFunc) (mc *MiddlewareController, err error) {
	var mo MiddlewareOptions = MiddlewareOptions{Logging: false}

	mc = &MiddlewareController{options: mo, AuthorizationFunction: nil}

	// loop through MiddlewareOptsFuncs passed in and process them.
	for _, fnc := range optsfuncs {
		fnc(&mo)
	}

	// set logging flag after option functions have been processed.
	mc.options.Logging = mo.Logging
	mc.AuthorizationFunction = mo.AuthorizationFunction

	return mc, nil
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

// function designed to set the authorization function of a MiddlewareOptions
// object. this will be assigned to the MiddlewareController that processes
// the MiddlewareOptions struct.
func WithAuthorization(af AuthFunc) MiddlewareOptsFunc {
	return func(mo *MiddlewareOptions) error {
		mo.AuthorizationFunction = af
		return nil
	}
}

// function designed to set the Logging flag of a MiddlewareOptions object.
// this will set the Logging variable to true.
func WithLogging(mo *MiddlewareOptions) (err error) {
	mo.Logging = true
	return nil
}

// function designed to handle writing a JSON payload to
// an HTTP response.
//
// if there is an issue encountered while encoding the
// JSON payload, the error will be returned, otherwise
// nil will be returned.
func WriteJSON(w *http.ResponseWriter, status int, v any) (err error) {
	(*w).Header().Set("Status-Code", fmt.Sprintf("%d", status))
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(status)
	return json.NewEncoder(*w).Encode(v)
}
