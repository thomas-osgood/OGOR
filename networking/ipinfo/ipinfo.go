package ipinfo

// package that holds objects and functions related to
// querying the site IPInfo.io and get information
// related to a given IP address. this package can be
// useful in target enumeration and analysis.
//
// information regarding IPInfo.io's API can be found at https://ipinfo.io/developers.

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// function designed to set the default configuration
// for an IPInfoQuery object.
func defaultOptions() (options ConfigStruct) {
	return ConfigStruct{Token: ""}
}

// function designed to set the token value for a
// given ConfigStruct. this should be passed into
// the NewQueryer function upon creation of a new
// IPInfoQuery object.
func WithToken(token string) ConfigFunc {
	var sanitizedToken string = token

	// remove any newlines from string
	sanitizedToken = strings.ReplaceAll(sanitizedToken, "\n", "")

	// remove any carriage-returns from string
	sanitizedToken = strings.ReplaceAll(sanitizedToken, "\r", "")

	if !strings.Contains(strings.ToLower(sanitizedToken), "bearer") {
		sanitizedToken = fmt.Sprintf("Bearer %s", sanitizedToken)
	}

	return func(cfg *ConfigStruct) {
		cfg.Token = sanitizedToken
	}
}

// function designed to create a new IPInfoQuery object and
// return it to the user.
func NewQueryer(options ...ConfigFunc) (IPQueryer *IPInfoQuery, err error) {
	var defaultConfig ConfigStruct = defaultOptions()

	// set user-defined configuration options if
	// any were passed in.
	for _, fn := range options {
		fn(&defaultConfig)
	}

	// build new IPInfoQuery object with correct
	// configuration options.
	IPQueryer = &IPInfoQuery{
		Options: defaultConfig,
	}

	return IPQueryer, nil
}

// function designed to reach out to ipinfo.io and
// return information related to the given IP address.
// this data will be returned in a struct that can
// be passed as JSON if desired. if something goes
// wrong during this function, an empty struct will
// be returned along with an error.
func (iq *IPInfoQuery) QueryAddress(ipaddress string) (info IPInfoStruct, err error) {
	var bodycontent []byte
	var client http.Client = http.Client{}
	var req *http.Request
	var resp *http.Response
	var targeturl string

	// make sure passed in IP address is not blank string.
	if len(ipaddress) < 1 {
		return IPInfoStruct{}, errors.New("ip address must be a non-zero length string")
	}

	// create URL to query and grab information from.
	targeturl = fmt.Sprintf("%s/%s/json", ipinfobase, ipaddress)

	// setup new request object. this is used rather than "http.Get"
	// so the user can set the authorization header if necessary.
	req, err = http.NewRequest(http.MethodGet, targeturl, nil)
	if err != nil {
		return IPInfoStruct{}, err
	}

	// if user specified an access token, add the
	// Authorization header.
	if len(iq.Options.Token) > 0 {
		req.Header.Set("Authorization", iq.Options.Token)
	}

	// make request to target URL.
	resp, err = client.Do(req)
	if err != nil {
		return IPInfoStruct{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return IPInfoStruct{}, errors.New(fmt.Sprintf("[query] unable to get IP information (%s)", resp.Status))
	}

	// read body of response.
	bodycontent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return IPInfoStruct{}, err
	}

	// load JSON from response into info struct.
	err = json.Unmarshal(bodycontent, &info)
	if err != nil {
		return IPInfoStruct{}, err
	}

	return info, nil
}

// function designed to determine whether the queried IP
// address is part of an internal network or open to the
// public. this is determined by checking the Bogon flag
// that gets returned by IPInfo.io.
func (inf *IPInfoStruct) IsInternalAddress() (internal bool, err error) {

	// make sure the struct has an IP Address.
	// if no IP Address present in the struct,
	// no query was made and there cannot be a
	// determination of Internal/Public.
	if len(inf.IP) < 1 {
		return false, errors.New("no IP address information. cannot conduct check.")
	}

	return inf.Bogon, nil
}
