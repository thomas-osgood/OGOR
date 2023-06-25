package apis

import (
	"errors"
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
	var testaddr string = strings.ToLower(ipaddr)

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
