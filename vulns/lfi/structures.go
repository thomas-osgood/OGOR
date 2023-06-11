package lfichecker

import "net/http"

// structure defining the LFIChecker object that will
// be used to check for LFI/Directory Traversal.
type LFIChecker struct {

	// respone length of a known bad route
	BadLength int

	// mapping of parameter bad values to return lengths.
	// these were the lengths returned when CheckBadLengthParams
	// was executed.
	BadLengthParams map[string]int

	// a route that will return a "404 Not Found" response. this route
	// will be used in various locations to check for LFI.
	BadRoute string

	// response length for a blank parameter.
	// this is only used when URL parameters are specified.
	BlankLength map[string]int

	// this is the HTTP client that will be conducting the
	// requests to the target.
	Checker LFIClient

	// LFI filter evasion techniques discovered. this will
	// be populated with successful evasion techniques when
	// checking for the LFI signature using CheckSignature.
	Evasions []string

	// response length of a known good route.
	GoodLength int

	// a route that will return a "200 OK" response. this route
	// will be used in various locations.
	GoodRoute string

	// LFI options associated with this checker
	Options LFIOptions

	// file to target when testing for LFI
	TargetFile string

	// return length of a target file test (no param).
	TestLength int

	// return length of a target file test using a parameter.
	TestLengthParams map[string]int

	// slice holding the vulnerable parameters that have been discovered.
	VulnerableParams map[string]string
}

// structure defining an LFIClient object that will be used
// to conduct requests to the target.
type LFIClient struct {

	// HTTP client that will be making requests.
	client http.Client

	// base url of target to check (eg: http://example.com). all
	// routes will be based off of this url.
	baseurl string
}

// structure defining the various LFI testing options the checker has.
type LFIOptions struct {

	// URL parameters to test when checking for LFI. if this slice
	// is empty, no parameters will be tested.
	//
	// default: empty
	Parameters map[string]string

	// switch indicating whether to use double URL encoding to attempt
	// to evade directory traversal filters.
	//
	// default: false
	DoubleEncoding bool

	// switch indicating whether to attempt to connect to the target
	// using HTTPS.
	//
	// default: false
	SSLConnection bool
}
