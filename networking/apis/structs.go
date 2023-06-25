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
