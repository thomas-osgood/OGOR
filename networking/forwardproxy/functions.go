package forwardproxy

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/thomas-osgood/OGOR/networking/validations"
)

// function designed to create and initialize a new
// proxy forwarder object. this will be used to proxy
// the traffic sent to the server.
func NewForwarder(userOptions ...ForwarderOptionsFunc) (forwarder *Forwarder, err error) {
	var defaultOptions ForwarderOptions = ForwarderOptions{}
	var defaultTransport http.RoundTripper = &http.Transport{IdleConnTimeout: DEFAULT_TIMEOUT}
	var fn ForwarderOptionsFunc

	// create instance of forwarder object.
	forwarder = new(Forwarder)

	defaultOptions.ForwardTransport = defaultTransport
	defaultOptions.Portno = DEFAULT_PORTNO
	defaultOptions.Server = &http.Server{}

	// loop through the user-provided settings and
	// set the appropriate variable. if any of the
	// settings errors out, return an error.
	for _, fn = range userOptions {
		err = fn(&defaultOptions)
		if err != nil {
			return nil, err
		}
	}

	defaultOptions.Server.Addr = fmt.Sprintf(":%d", defaultOptions.Portno)
	defaultOptions.Server.TLSConfig = &tls.Config{
		MinVersion:               tls.VersionTLS13,
		PreferServerCipherSuites: true,
	}

	//////////////////////////////////////////////////////////////
	// set the values for the object that is going to
	// be returned and used.
	//////////////////////////////////////////////////////////////

	forwarder.forwardTransport = defaultOptions.ForwardTransport
	forwarder.portno = defaultOptions.Portno
	forwarder.server = defaultOptions.Server

	return forwarder, nil
}

// function designed to set the serve port for the forwarder.
// this can be passed to the NewForwarder function during
// creation of a new object.
func UseServePort(portno int) ForwarderOptionsFunc {
	return func(fo *ForwarderOptions) (err error) {
		err = validations.ValidateNetworkPort(portno)
		if err != nil {
			return err
		}

		fo.Portno = portno

		return nil
	}
}

// function designed to set the API Middlware Controller
// for the forwarder. this can be passed to the NewForwarder
// function during creation of a new object.
func UseTransport(transport http.RoundTripper) ForwarderOptionsFunc {
	return func(fo *ForwarderOptions) error {
		fo.ForwardTransport = transport
		return nil
	}
}
