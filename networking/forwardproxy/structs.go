package forwardproxy

import (
	"net/http"
)

// struct defining the forwarder object used
// to proxy traffic.
type Forwarder struct {
	// transport to use to transmit the requests.
	forwardTransport http.RoundTripper
	portno           int
	server           *http.Server
}

// struct defining various options the user can
// set when initializing a new forwarder object.
type ForwarderOptions struct {
	// transport to use to transmit the requests.
	ForwardTransport http.RoundTripper
	// port to host the forwarder on.
	Portno int
	// server to use to server the forwarder.
	Server *http.Server
}
