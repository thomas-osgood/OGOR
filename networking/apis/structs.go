package apis

// struct designed to represent a JSON error return that can
// be used to deliver information about an error that has been
// thrown during execution of an API/Server function.
type ErrorStruct struct {
	// the code associated with the given error.
	ErrorCode int `json:"code"`

	// a message detailing the error that has been thrown.
	ErrorMessage string `json:"message"`
}

// structure designed to represent a middleware controller that
// can be used to filter requests to the API. this will include
// various generic settings that can be controlled by the user.
type MiddlewareController struct {

	// slice holding IP addresses that are blacklisted. these
	// can be used to restrict who can contact the API.
	AddressBlacklist []string

	// slice holding all the headers reuqired for a request to
	// be properly handled by the API. this is meant to be a list
	// that all endpoints require and not meant to be specific
	// to a single endpoint.
	RequiredHeaders []string
}
