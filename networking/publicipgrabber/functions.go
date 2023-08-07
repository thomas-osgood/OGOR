package publicipgrabber

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// function designed to create and initialize a new
// PubliIPGrabber object. the user can pass in option
// functions to change the configuration.
func NewPublicIPGrabber(optfuncs ...PublicIPGrabberOptFunc) (grabber *PublicIPGrabber, err error) {
	var fn PublicIPGrabberOptFunc
	var options PublicIPGrabberOptions = PublicIPGrabberOptions{Client: http.DefaultClient}

	grabber = &PublicIPGrabber{}

	// loop through and set configuration options.
	for _, fn = range optfuncs {
		err = fn(&options)
		if err != nil {
			return nil, err
		}
	}

	grabber.client = options.Client

	return grabber, nil
}

// function designed to contact the ipify api to pull down
// the current machine's public IPv4 address. this returns
// only the IP address and no additional information. this
// can be used to quickly query the public IPv4 without the
// need to create a PublicIPGrabber object.
//
// for further info: https://www.ipify.org
func QueryIPIfy() (publicip string, err error) {
	var bodycontent []byte
	var client *http.Client = http.DefaultClient
	var resp *http.Response
	const targeturl string = "https://api.ipify.org"

	resp, err = client.Get(targeturl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return "", errors.New(fmt.Sprintf("error contacting ipify: %s", resp.Status))
	}

	bodycontent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if len(bodycontent) < 1 {
		return "", errors.New("no ip address returned")
	}

	publicip = strings.TrimSpace(string(bodycontent))

	return publicip, nil
}

// function designed to contact the ipify api to pull down
// the current machine's public IPv6 address. this returns
// only the IP address and no additional information. this
// can be used to quickly query the public IPv6 without the
// need to create a PublicIPGrabber object.
//
// note: if no IPv6 address is found, the IPv4 address will
// be returned.
//
// for further info: https://www.ipify.org
func QueryIPIfy6() (publicip string, err error) {
	var bodycontent []byte
	var client *http.Client = http.DefaultClient
	var resp *http.Response
	const targeturl string = "https://api64.ipify.org"

	resp, err = client.Get(targeturl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return "", errors.New(fmt.Sprintf("error contacting ipify: %s", resp.Status))
	}

	bodycontent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if len(bodycontent) < 1 {
		return "", errors.New("no ipv6 address returned")
	}

	publicip = strings.TrimSpace(string(bodycontent))

	return publicip, nil
}

// function designed to set the PublicIPGrabberOptions client.
func WithClient(client *http.Client) PublicIPGrabberOptFunc {
	return func(pio *PublicIPGrabberOptions) error {
		pio.Client = client
		return nil
	}
}
