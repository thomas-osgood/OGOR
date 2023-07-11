package proxyscrape

import (
	"errors"
	"fmt"
	"strings"
)

// function designed to create, initialize and return a new
// ProxyScraper instance that can be used to contact the API.
func NewProxyScraper(opts ...OptsFunc) (scraper *ProxyScraper, err error) {
	var options ScraperOptions = ScraperOptions{}
	scraper = &ProxyScraper{Proxies: ProxyList{}}

	options.Anonymity = ANONYMITY_ELITE
	options.Country = "us"
	options.Protocol = PROTO_HTTP
	options.SSL = HTTPS_ALL
	options.Timeout = MAX_TIMEOUT

	for _, fn := range opts {
		err = fn(&options)
		if err != nil {
			return nil, err
		}
	}

	scraper.anonymity = options.Anonymity
	scraper.country = options.Country
	scraper.protocol = options.Protocol
	scraper.ssl = options.SSL
	scraper.timeout = options.Timeout

	return scraper, nil
}

// function designed to set the anonymity option for the
// current ScraperOptions struct.
func UsingAnonymity(anontype string) OptsFunc {
	return func(so *ScraperOptions) error {
		anontype = strings.ToLower(anontype)
		switch anontype {
		case "all":
			so.Anonymity = ANONYMITY_ALL
		case "transparent":
			so.Anonymity = ANONYMITY_TRANSPARENT
		case "anonymous":
			so.Anonymity = ANONYMITY_ANONYMOUS
		case "elite":
			so.Anonymity = ANONYMITY_ELITE
		default:
			return errors.New("invalid anonymity type")
		}
		return nil
	}
}

// function designed to set the country option for the
// current ScraperOptions struct.
func UsingCountry(country string) OptsFunc {
	return func(so *ScraperOptions) error {
		country = strings.ToLower(country)
		if (country != "all") && (len(country) != 2) {
			return errors.New("country code must be two letters if it is not \"all\"")
		}
		so.Country = country
		return nil
	}
}

// function designed to set the protocol option for the
// current ScraperOptions struct.
func UsingProtocol(protocol string) OptsFunc {
	return func(so *ScraperOptions) error {
		protocol = strings.ToLower(protocol)
		switch protocol {
		case "all":
			so.Protocol = PROTO_ALL
		case "http":
			so.Protocol = PROTO_HTTP
		case "socks4":
			so.Protocol = PROTO_SOCKS4
		case "socks5":
			so.Protocol = PROTO_SOCKS5
		default:
			return errors.New(fmt.Sprintf("invalid protocol \"%s\"", protocol))
		}
		return nil
	}
}

// function designed to set the ssl option for the
// current ScraperOptions struct.
func UsingSSL(ssltype string) OptsFunc {
	return func(so *ScraperOptions) error {
		ssltype = strings.ToLower(ssltype)
		switch ssltype {
		case "all":
			so.SSL = HTTPS_ALL
		case "yes":
			so.SSL = HTTPS_YES
		case "no":
			so.SSL = HTTPS_NO
		default:
			return errors.New("invalid SSL option.")
		}
		return nil
	}
}

// function designed to set the timeout option for the
// current ScraperOptions struct.
func UsingTimeout(timeout int) OptsFunc {
	return func(so *ScraperOptions) (err error) {
		err = ValidateTimeout(timeout)
		if err != nil {
			return err
		}
		so.Timeout = timeout
		return nil
	}
}

// function designed to validate the timeout value for the ProxyScraper.
func ValidateTimeout(timeout int) (err error) {
	if (timeout < 1) || (timeout > MAX_TIMEOUT) {
		return errors.New(fmt.Sprintf("timeout must be in range 1 - %d", MAX_TIMEOUT))
	}
	return nil
}
