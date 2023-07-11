// package designed to contact proxyscrape's API and pull down a list
// of proxies based on the options the user selects.
//
// API documentation:
//
//	https://docs.proxyscrape.com/#1ec9e5ed-0dce-4511-91e1-ebe99f7bd88d
package proxyscrape

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// function designed to contact the ProxyScraper API and
// pull down a list of proxies.
func (p *ProxyScraper) GetProxies() (err error) {
	var anonymity string
	var bodycontent []byte
	var bodystring string
	var client http.Client = http.Client{Timeout: time.Duration(10) * time.Second}
	var req *http.Request
	var resp *http.Response
	var params url.Values = url.Values{}
	var protocol string
	var ssl string

	anonymity, err = p.getAnonymity()
	if err != nil {
		return err
	}

	protocol, err = p.getProtocolString()
	if err != nil {
		return err
	}

	ssl, err = p.getSSLOption()
	if err != nil {
		return err
	}

	// set URL parameters
	params.Set("anonymity", anonymity)
	params.Set("request", "displayproxies")
	params.Set("protocol", protocol)
	params.Set("country", p.country)
	params.Set("ssl", ssl)
	params.Set("timeout", p.getTimeoutString())

	// create request
	req, err = http.NewRequest(http.MethodGet, BASE_URL, nil)
	if err != nil {
		return err
	}
	req.URL.RawQuery = params.Encode()

	// make request
	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// call returned with status of 400 or greater.
	if resp.StatusCode >= http.StatusBadRequest {
		return errors.New(fmt.Sprintf("proxyscraper api unreachable (%s)", resp.Status))
	}

	// get return body in bytes.
	bodycontent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// convert bytes to string, ignoring blank line at ent.
	bodystring = strings.Trim(string(bodycontent), " \n\r")

	// split return string by newline and make slice.
	p.Proxies.Proxies = strings.Split(bodystring, "\n")

	return nil
}

// function designed to get the anonymity level specified
// for the ProxyScraper.
func (p *ProxyScraper) getAnonymity() (anonymity string, err error) {

	if p.anonymity == ANONYMITY_ALL {
		anonymity = "all"
	} else if p.anonymity == ANONYMITY_TRANSPARENT {
		anonymity = "transparent"
	} else if p.anonymity == ANONYMITY_ANONYMOUS {
		anonymity = "anonymous"
	} else if p.anonymity == ANONYMITY_ELITE {
		anonymity = "elite"
	} else {
		return "", errors.New("invalid anonymity specified.")
	}

	return anonymity, nil
}

// function designed to get the string representation of the
// protocol specified for the ProxyScraper.
func (p *ProxyScraper) getProtocolString() (protocol string, err error) {
	switch p.protocol {
	case PROTO_ALL:
		protocol = "all"
	case PROTO_HTTP:
		protocol = "http"
	case PROTO_SOCKS4:
		protocol = "socks4"
	case PROTO_SOCKS5:
		protocol = "socks5"
	default:
		return "", errors.New("invalid protocol option")
	}
	return protocol, nil
}

// function designed to get the string representation of the
// SSL option specified for the ProxyScraper.
func (p *ProxyScraper) getSSLOption() (ssl string, err error) {
	switch p.ssl {
	case HTTPS_ALL:
		ssl = "all"
	case HTTPS_YES:
		ssl = "yes"
	case HTTPS_NO:
		ssl = "no"
	default:
		return "", errors.New(fmt.Sprintf("invalid option for SSL. expecting %d, %d, or %d.", HTTPS_ALL, HTTPS_YES, HTTPS_NO))
	}
	return ssl, nil
}

// function designed to return the string representation of the
// specified timeout.
func (p *ProxyScraper) getTimeoutString() string {
	return fmt.Sprintf("%d", p.timeout)
}
