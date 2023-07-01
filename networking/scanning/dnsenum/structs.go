package dnsenum

import "net/http"

// structure designed to represent a DNS enumerator.
// this will hold all the necessary information for
// the user to execute DNS enumeration.
type Enumerator struct {

	// HTTP client that will make a request to the
	// top-level domain and test for subdomains.
	Client *http.Client

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
}

type EnumOpts struct {

	// specified if the user wants the enumerator to
	// use an existing HTTP client. this will override
	// the client defined in the Enuerator struct.
	ExistingClient *http.Client

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
