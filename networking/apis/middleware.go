package apis

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/netip"
	"strings"
)

// function designed to check if a given IP address string is in the
// AddressBlacklist slice for the middleware. if the blacklist contains
// the given IP address nil will be returned, otherwise an error
// will be returned.
//
// the address comparison is case-insensitive.
func (mc *MiddlewareController) Blacklisted(ipaddr string) (err error) {
	var found bool = false
	var netaddr netip.Addr
	var testaddr string

	testaddr, _, err = net.SplitHostPort(ipaddr)
	if err != nil {
		return err
	}

	netaddr, err = netip.ParseAddr(testaddr)
	if err != nil {
		return err
	}
	testaddr = netaddr.String()

	// loop through blacklist and check given address against each value.
	for _, badaddr := range mc.AddressBlacklist {
		badaddr = strings.ToLower(badaddr)
		if testaddr == badaddr {
			found = true
			break
		}
	}

	// found flag is still false. the address is not in the blacklist.
	if !found {
		return errors.New("address not found in blacklist.")
	}

	return nil
}

// function designed to more elegantly handle HTTP routing functions.
// this is, essentially, a middleware controller that extends the
// HandleFunc function. this takes an APIFunc as an argument and
// processes it.
func (mc *MiddlewareController) MakeHTTPHandleFunc(fnc APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		// make sure the address making the request is not in the blacklist.
		err = mc.Blacklisted(r.RemoteAddr)
		if err == nil {
			log.Printf("denied blacklisted address: \"%s\"\n", r.RemoteAddr)
			log.Printf("\tmethod: \"%s\"\n", r.Method)
			log.Printf("\turi: \"%s\"\n", r.RequestURI)
			return
		} else if err.Error() != "address not found in blacklist." {
			log.Printf("error checking blacklist: %s\n", err.Error())
			log.Printf("denying request...\n")
			return
		}

		// make sure the user requesting the endpoint is authorized.
		if mc.AuthorizationFunction != nil {
			err = mc.AuthorizationFunction(r)
			if err != nil {
				log.Printf("unautorized request from \"%s\" to \"%s\" blocked\n", r.RemoteAddr, r.RequestURI)
				err = ReturnErrorJSON(&w, http.StatusUnauthorized, "unauthorized")
				if err != nil {
					w.Header().Set("Content-Type", "text/plain")
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("unauthorized"))
				}
				return
			}
		}

		// if logging is turned on, log the current request. this
		// prints out the Method, Remote Address, and URL.
		if mc.options.Logging {
			log.Printf("\"%s\" request from \"%s\" to \"%s\"\n", r.Method, r.RemoteAddr, r.URL.RequestURI())
		}

		// process request
		err = fnc(w, r)

		// there was an error processing the request. return a plaint text
		// response showing the error.
		if err != nil {
			w.Header().Set("Status-Code", fmt.Sprintf("%d", http.StatusBadRequest))
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
