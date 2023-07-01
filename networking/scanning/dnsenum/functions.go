package dnsenum

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// function designed to create an initialize a DNS
// enumerator object.
func NewEnumerator(tld string, opts ...EnumOptsFunc) (enumerator *Enumerator, err error) {
	if len(tld) < 1 {
		return nil, errors.New("invalid TLD passed in")
	}

	var options EnumOpts = EnumOpts{ExistingClient: nil, TestHeader: false, Wordlist: "subdomains.txt", Timeout: 10}

	enumerator = &Enumerator{TLD: tld}

	// loop through EnumOptsFuncs passed in and
	// set user-defined values.
	for _, fn := range opts {
		err = fn(&options)
		if err != nil {
			fmt.Printf("[init] error executing \"%v\"\n", fn)
		}
	}

	if options.ExistingClient != nil {
		enumerator.Client = options.ExistingClient
	} else {
		enumerator.Client = &http.Client{}
	}

	enumerator.Client.Timeout = time.Duration(options.Timeout * float64(time.Second))

	enumerator.TestHeader = options.TestHeader
	enumerator.Wordlist = options.Wordlist

	return enumerator, nil
}

// opts func to set the TestHeader flag, indicating to place
// the subdomain value in the Host header during enumeration.
func UseHeader(eo *EnumOpts) error {
	eo.TestHeader = true
	return nil
}

// opts func to specify the timeout duration for the client.
func WithTimeout(duration float64) EnumOptsFunc {
	return func(eo *EnumOpts) error {
		eo.Timeout = duration
		return nil
	}
}

// opts func to specify the wordlist to use during enumeration.
func WithWordlist(wordlist string) EnumOptsFunc {
	return func(eo *EnumOpts) error {
		eo.Wordlist = wordlist
		return nil
	}
}
