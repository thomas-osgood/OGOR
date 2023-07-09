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

// function designed to log a middleware event.
func (mc *MiddlewareController) LogEvent(message string, severity int) {

	// set text color based on the severity of the event if
	// coloring flag is set.
	if mc.options.Coloring {
		switch severity {
		case EVENT_ERROR:
			message = mc.formatter.RedText(message)
		case EVENT_SUCCESS:
			message = mc.formatter.GreenText(message)
		case EVENT_WARNING:
			message = mc.formatter.YellowText(message)
		default:
			message = mc.formatter.BlueText(message)
		}
	}

	// output event message.
	log.Printf("%s\n", message)

	return
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
			mc.LogEvent(fmt.Sprintf("denied blacklisted address: \"%s\"", r.RemoteAddr), EVENT_ERROR)
			mc.LogEvent(fmt.Sprintf("\tmethod: \"%s\"", r.Method), EVENT_ERROR)
			mc.LogEvent(fmt.Sprintf("\turi: \"%s\"", r.RequestURI), EVENT_ERROR)
			return
		} else if err.Error() != "address not found in blacklist." {
			mc.LogEvent(fmt.Sprintf("error checking blacklist: %s", err.Error()), EVENT_ERROR)
			mc.LogEvent("denying request...", EVENT_ERROR)
			return
		}

		// make sure the user requesting the endpoint is authorized.
		if mc.AuthorizationFunction != nil {
			err = mc.AuthorizationFunction(r)
			if err != nil {
				mc.LogEvent(fmt.Sprintf("unauthorized request from \"%s\" to \"%s\" blocked", r.RemoteAddr, r.RequestURI), EVENT_ERROR)
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
			mc.LogEvent(fmt.Sprintf("\"%s\" request from \"%s\" to \"%s\"", r.Method, r.RemoteAddr, r.URL.RequestURI()), EVENT_INFO)
		}

		// process request
		err = fnc(w, r)

		// there was an error processing the request. return a plaint text
		// response showing the error.
		if err != nil {
			if mc.options.Logging {
				mc.LogEvent(fmt.Sprintf("error processing request: %s", err.Error()), EVENT_ERROR)
			}

			w.Header().Set("Status-Code", fmt.Sprintf("%d", http.StatusBadRequest))
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
