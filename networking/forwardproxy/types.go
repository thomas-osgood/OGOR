package forwardproxy

// type definition explaining a Forwarder Options
// function. this will be used during initialization
// of a new Forwarder object to set user-controlled
// options and variables.
type ForwarderOptionsFunc func(*ForwarderOptions) error
