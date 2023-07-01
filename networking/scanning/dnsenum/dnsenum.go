package dnsenum

import (
	"fmt"
	"net/http"
)

// function designed to see if the TLD provided is reachable
// using HEAD and is, therefore, a valid site or if the domain
// is comprised enitrely of subdomains (if any).
func (e *Enumerator) TestTLD() (err error) {
	err = e.TestSubdomainHead(e.TLD, false)
	if err != nil {
		err = e.TestSubdomainHead(e.TLD, true)
		if err != nil {
			return err
		}
	}
	return nil
}

// function designed to test a subdomain to see if returns without
// error when a HEAD request is made against it.
//
// note: the subdomain must be <subdomain>.<tld> (ex: test.example.com)
func (e *Enumerator) TestSubdomainHead(subdomain string, https bool) (err error) {
	var resp *http.Response
	var targeturl string

	if https {
		targeturl = fmt.Sprintf("https://%s", subdomain)
	} else {
		targeturl = fmt.Sprintf("http://%s", subdomain)
	}

	resp, err = e.Client.Head(targeturl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
