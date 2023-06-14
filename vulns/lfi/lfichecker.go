// this package is designed to assist a pentester or ethical hacker on
// checking for Local File Inclusion (LFI) or Directory Traversal
// vulnerabilities in a target site.
package lfichecker

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/thomas-osgood/OGOR/misc/generators"
)

// function designed to check an individual URL parameter for
// an LFI/Directory Traversal vulnerability.
func (l *LFIChecker) CheckParameter(param string) (err error) {
	var bodycontent []byte
	var bodylen int
	var detected bool = false
	var params url.Values = url.Values{}
	var req *http.Request
	var resp *http.Response
	var targeturl string = fmt.Sprintf("%s/%s", l.Checker.baseurl, l.GoodRoute)
	const traverselen int = 10
	var traversestr string = ""

	for _, bypass := range lfipatterns {
		for i := 0; i < traverselen; i++ {
			traversestr = fmt.Sprintf("%s%s", traversestr, bypass)
		}

		// set all parameter values
		for k, v := range l.Options.Parameters {
			params.Set(k, v)
		}

		// poison parameter value to test
		params.Set(param, fmt.Sprintf("%s%s", traversestr, l.Options.TargetFile))

		req, err = http.NewRequest(http.MethodGet, targeturl, nil)
		if err != nil {
			return err
		}

		req.URL.RawQuery = params.Encode()

		resp, err = l.Checker.client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 400 {
			return errors.New(fmt.Sprintf("bad status code with \"%s\" (%s)", param, resp.Status))
		}

		bodycontent, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		bodylen = len(bodycontent)

		if (bodylen == l.BadLength) || (bodylen == l.GoodLength) {
			continue
		}

		l.VulnerableParams[param] = bypass
		detected = true
		break
	}

	if !detected {
		return errors.New(fmt.Sprintf("lfi not detected in \"%s\"", param))
	}

	return nil
}

// function designed to check for an LFI signature using the current
// LFIChecker configuration. this will compare the various lengths and
// attempt to determine if LFI is present on the target. if no LFI
// is present, an error will be returned.
func (l *LFIChecker) CheckSignature() (err error) {
	var bclen int
	var bodycontent []byte
	var bypass string
	var goodbadcheck bool
	var patternstring string
	var randfolder string

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

	// generate random folder name to use for LFI signature test.
	randfolder, err = generators.GenRandomName(generators.DEFAULT_RAND_MIN, generators.DEFAULT_RAND_MAX)
	if err != nil {
		return err
	}

	for _, bypass = range lfipatterns {
		patternstring = ""
		for i := 0; i < 10; i++ {
			patternstring = fmt.Sprintf("%s%s", patternstring, bypass)
		}
		patternstring = fmt.Sprintf("%s%s", patternstring, randfolder)
		bodycontent, err = l.GetBodyContent(patternstring)
		if err != nil {
			continue
		}
		bclen = len(bodycontent)

		// make sure the returned length does not equal the
		// good response length and the bad response length.
		// if neither is a match, LFI is likely.
		if (bclen != l.GoodLength) && (bclen != l.BadLength) {
			l.Evasions = append(l.Evasions, bypass)
		}
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

	// set the blank length variable to check.
	err = l.GetBlankLength()
	if err != nil {
		return err
	}

	for k := range l.Options.Parameters {
		if (l.GoodLength == l.BadLengthParams[k]) || (l.GoodLength == l.BlankLength[k]) {
			continue
		}

		if err = l.CheckParameter(k); err != nil {
			continue
		}
	}

	if len(l.VulnerableParams) < 1 {
		return errors.New("no vulnerable parameters discovered")
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

// function designed to check for the return length when a bad parameter
// value is passed in.
func (l *LFIChecker) GetBadLengthParams() (err error) {
	var bodycontent []byte
	var bodylen int
	var randval string
	var req *http.Request
	var resp *http.Response
	var targeturl string = fmt.Sprintf("%s/%s", l.Checker.baseurl, l.GoodRoute)
	var params url.Values = url.Values{}

	if len(l.Options.Parameters) < 1 {
		return errors.New("no parameters to test")
	}

	randval, err = generators.GenRandomName(generators.DEFAULT_RAND_MIN, generators.DEFAULT_RAND_MAX)
	if err != nil {
		return err
	}

	// setup HTTP request to target.
	req, err = http.NewRequest(http.MethodGet, targeturl, nil)
	if err != nil {
		return err
	}

	// loop through given parameters and record the bad return length of each.
	for curkey := range l.Options.Parameters {
		for k, v := range l.Options.Parameters {
			params.Set(k, v)
		}

		params.Set(l.Options.Parameters[curkey], randval)

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
		l.BadLengthParams[curkey] = bodylen
	}

	return nil
}

// function designed to get the return length when a blank parameter is passed
// to the target.
func (l *LFIChecker) GetBlankLength() (err error) {
	var bodycontent []byte
	var bodylen int
	var req *http.Request
	var resp *http.Response
	var targeturl string = fmt.Sprintf("%s/%s", l.Checker.baseurl, l.GoodRoute)
	var params url.Values = url.Values{}

	if len(l.Options.Parameters) < 1 {
		return errors.New("no parameters to test")
	}

	// setup HTTP request to target.
	req, err = http.NewRequest(http.MethodGet, targeturl, nil)
	if err != nil {
		return err
	}

	for curkey := range l.Options.Parameters {
		for k, v := range l.Options.Parameters {
			params.Set(k, v)
		}

		keys := make([]string, 0, len(l.Options.Parameters))
		for k := range l.Options.Parameters {
			keys = append(keys, k)
		}

		params.Set(l.Options.Parameters[curkey], "")

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
		l.BlankLength[curkey] = bodylen
	}

	return nil
}

// function designed to perform an HTTP GET request on a target
// route, returning the body content of the resonse.  this is
// useful when checking the response length of a target route.
func (l *LFIChecker) GetBodyContent(route string) (bodycontent []byte, err error) {
	var targetURL string = fmt.Sprintf("%s/%s", l.Checker.baseurl, route)

	resp, err := l.Checker.client.Get(targetURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodycontent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodycontent, nil
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

	// if user specified URL parameters, add them to the request.
	if len(l.Options.Parameters) > 0 {
		var params url.Values = url.Values{}

		for k, v := range l.Options.Parameters {
			params.Set(k, v)
		}

		req.URL.RawQuery = params.Encode()
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

// private function designed to analyze the Checker's base url and
// attempt to determine the file extension/type of the given file.
// this can help determine if PHP filters should be tested, etc.
func (l *LFIChecker) getfiletype() (filetype string, err error) {
	var baselen int
	var baseurl string
	var splitlen int
	var splitstring []string = []string{}

	baseurl = l.Checker.baseurl
	baselen = len(baseurl)
	if baseurl[baselen-1] == '/' {
		baseurl = baseurl[:baselen-1]
	}

	splitstring = strings.Split(baseurl, "?")
	splitstring = strings.Split(splitstring[0], "/")

	splitlen = len(splitstring)

	splitstring = strings.Split(splitstring[splitlen-1], ".")

	splitlen = len(splitstring)

	// if the string split by "." does not have a length of 2 or greater,
	// no file extension is present in the filename (eg: /index) and the
	// file type cannot be determined using this function.
	if splitlen < 2 {
		return "", errors.New("page name does not have an extension")
	}

	// splitstring's last index will hold the file extension.
	filetype = splitstring[splitlen-1]

	return filetype, nil
}

// function designed to set the GoodRoute parameter in the LFIChecker object.
func (l *LFIChecker) SetGoodRoute(route string) (err error) {
	l.GoodRoute = route
	return nil
}
