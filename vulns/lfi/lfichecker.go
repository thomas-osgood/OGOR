// this package is designed to assist a pentester or ethical hacker on
// checking for Local File Inclusion (LFI) or Directory Traversal
// vulnerabilities in a target site.
package lfichecker

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// function designed to check for an LFI signature using the current
// LFIChecker configuration. this will compare the various lengths and
// attempt to determine if LFI is present on the target. if no LFI
// is present, an error will be returned.
func (l *LFIChecker) CheckSignature() (err error) {
	var goodbadcheck bool

	// set the good length variable to check.
	err = l.GetGoodLength()
	if err != nil {
		return err
	}

	// set the bad length variable to check.
	err = l.GetBadLength()
	if err != nil {
		return err
	}

	// if good route content size is the same as bad route content size,
	// this check will fail.
	goodbadcheck = (l.BadLength != l.GoodLength)

	if !goodbadcheck {
		return errors.New("LFI signature not found on the target.")
	}
	return nil
}

// function designed to check for an LFI signature using the current
// LFIChecker configuration. this will target URL parameters, compare
// various lengths and attempt to determine if LFI is present on
// the target. if no LFI is present, an error will be returned.
func (l *LFIChecker) CheckSignatureWithParams() (err error) {

	// make sure there are URL parameters to check.
	if len(l.Options.Parameters) < 1 {
		return errors.New("no parameters to check")
	}

	return nil
}

// function designed to contact the target and get the
// length of a request that returns a 404 NOT FOUND response. this
// length can be used as part of the check for LFI/Directory Traversal.
func (l *LFIChecker) GetBadLength() (err error) {
	var bodycontent []byte
	var bodylen int
	var req *http.Request
	var resp *http.Response
	var targeturl string = fmt.Sprintf("%s/%s", l.Checker.baseurl, l.BadRoute)

	// setup HTTP request to target.
	req, err = http.NewRequest(http.MethodGet, targeturl, nil)
	if err != nil {
		return err
	}

	// make request to target.
	resp, err = l.Checker.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// read HTTP response.
	//
	// note, this does not check the return code because some sites
	// do not return 404 Not Found when responding with an error message.
	bodycontent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// get response length.
	bodylen = len(bodycontent)

	// set BadLength variable.
	l.BadLength = bodylen

	return nil
}

// function designed to contact the target and get the
// length of a request that returns a 200 OK response. this
// length can be used as part of the check for LFI/Directory Traversal.
func (l *LFIChecker) GetGoodLength() (err error) {
	var bodycontent []byte
	var bodylen int
	var req *http.Request
	var resp *http.Response
	var targeturl string = fmt.Sprintf("%s/%s", l.Checker.baseurl, l.GoodRoute)

	// setup HTTP request to target.
	req, err = http.NewRequest(http.MethodGet, targeturl, nil)
	if err != nil {
		return err
	}

	// make request to target.
	resp, err = l.Checker.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// if "GoodRoute" returns >= 400, return an error.
	if resp.StatusCode >= http.StatusBadRequest {
		return errors.New(fmt.Sprintf("Good Route is returning a non-ok status code (%s)", resp.Status))
	}

	// read HTTP response.
	bodycontent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// get response length.
	bodylen = len(bodycontent)

	// set GoodLength variable.
	l.GoodLength = bodylen

	return nil
}

// function designed to set the BadRoute parameter in the LFIChecker object.
func (l *LFIChecker) SetBadRoute(route string) (err error) {
	l.BadRoute = route
	return nil
}

// function designed to set the GoodRoute parameter in the LFIChecker object.
func (l *LFIChecker) SetGoodRoute(route string) (err error) {
	l.GoodRoute = route
	return nil
}
