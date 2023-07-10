package dnsenum

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/thomas-osgood/OGOR/output"
)

// function designed to create an initialize a DNS
// enumerator object.
func NewEnumerator(tld string, opts ...EnumOptsFunc) (enumerator *Enumerator, err error) {
	if len(tld) < 1 {
		return nil, errors.New("invalid TLD passed in")
	}

	var options EnumOpts = EnumOpts{ExistingClient: nil, TestHeader: false, Wordlist: "subdomains.txt", Timeout: 10, Https: false, Delay: 0, Display: false}

	enumerator = &Enumerator{TLD: tld, Discovered: []string{}}

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

	enumerator.delay = options.Delay
	enumerator.display = options.Display
	enumerator.https = options.Https
	enumerator.threads = options.ThreadCount

	enumerator.printer, err = output.NewOutputter()
	if err != nil {
		return nil, err
	}

	return enumerator, nil
}

// function designed to set the display flag, indicating to
// show the enumeration results as it is happening. this will
// display output using the enumerator's Outputter object.
func ShowOutput(eo *EnumOpts) error {
	eo.Display = true
	return nil
}

// opts func to set the TestHeader flag, indicating to place
// the subdomain value in the Host header during enumeration.
func UseHeader(eo *EnumOpts) error {
	eo.TestHeader = true
	return nil
}

// opts func to set the max delay between requests. this is
// to make sure the enumerator does not overload the target
// or get blocked.
func WithDelay(delay int) EnumOptsFunc {
	return func(eo *EnumOpts) error {
		if delay < 0 {
			return errors.New("delay must be >= 0")
		}
		eo.Delay = delay
		return nil
	}
}

// opts func to set the HTTPS flag, indicating to use HTTPS
// when enumerating subdomains for the target.
func WithHTTPS(eo *EnumOpts) error {
	eo.Https = true
	return nil
}

// opts func to specify the number of threads to use.
func WithThreadCount(count int) EnumOptsFunc {
	return func(eo *EnumOpts) error {
		if (count < 1) || (count > 100) {
			return errors.New("threadcount must be in range 1 - 100")
		}
		eo.ThreadCount = count
		return nil
	}
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
