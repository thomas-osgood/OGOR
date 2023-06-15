// package designed to be used to fingerprint websites,
// apis, servers, etc.
//
// when used in conjunction with other utilities and information,
// this can be very useful in a penetration test.
package fingerprinter

// structure defining the fingerprinter object. this
// will hold the target site and any other information
// necessary to fingerprint it.
type Fingerprinter struct {
	// public variables. these can be directly interacted with
	// by the person importing/using this module.

	Target string

	// private variables. these cannot be directly interacted
	// with by the person importing/using this module.

	allowedmethods []string
}
