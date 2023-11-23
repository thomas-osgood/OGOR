package forwardproxy

import (
	"net/http"
)

// struct defining the forwarder object used
// to proxy traffic.
type Forwarder struct {
	// transport to use to transmit the requests.
	forwardTransport http.RoundTripper
	logging          bool
	portno           int
	server           *http.Server
}

// struct defining various options the user can
// set when initializing a new forwarder object.
type ForwarderOptions struct {
	// transport to use to transmit the requests.
	ForwardTransport http.RoundTripper
	// flag indicating whether to log traffic to STDOUT.
	// this flag is set by default.
	Logging bool
	// port to host the forwarder on.
	Portno int
	// server to use to serve the forwarder.
	Server *http.Server
}
