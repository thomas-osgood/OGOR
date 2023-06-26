package apis

import "net/http"

// alias of http request function that returns an error. this will
// be handled by middleware and is used for more control/neater code.
type APIFunc func(http.ResponseWriter, *http.Request) error

// alias of a function that is used to manipulate the various parts
// of a MiddlewareOptions object.
type MiddlewareOptsFunc func(*MiddlewareOptions) error
