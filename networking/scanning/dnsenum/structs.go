package dnsenum

import (
	"net/http"

	"github.com/thomas-osgood/OGOR/output"
)

// structure designed to represent a DNS enumerator.
// this will hold all the necessary information for
// the user to execute DNS enumeration.
type Enumerator struct {

	// HTTP client that will make a request to the
	// top-level domain and test for subdomains.
	Client *http.Client

	// slice containing the discovered subdomains.
	Discovered []string

	// flag indicating whether the enumerator should test
	// the domain using "Host: <subdomain>" to search
	// for subdomains.
	TestHeader bool

	// top-level domain to test for subdomains.
	TLD string

	// length of top-level domain return.
	TLDLength int

	// wordlist to use for enumeration. this defaults
	// to subdomains.txt.
	Wordlist string

	// max delay time each worker thread should wait
	// in between requests. defaults to 0.
	delay int

	// specify HTTPS for testing.
	https bool

	// object used to output data.
	printer *output.Outputter
}

type EnumOpts struct {

	// max number of milliseconds a thread should wait
	// in between requests.
	Delay int

	// specified if the user wants the enumerator to
	// use an existing HTTP client. this will override
	// the client defined in the Enuerator struct.
	ExistingClient *http.Client

	// specify HTTPS for testing
	Https bool

	// flag indicating whether the enumerator should test
	// the domain using "Host: <subdomain>" to search
	// for subdomains.
	TestHeader bool

	// amount of time (in seconds) to wait before dropping
	// the request if it has not responded.
	Timeout float64

	// wordlist to use for enumeration. this defaults
	// to subdomains.txt.
	Wordlist string
}
