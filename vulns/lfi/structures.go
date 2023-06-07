package lfichecker

import "net/http"

// structure defining the LFIChecker object that will
// be used to check for LFI/Directory Traversal.
type LFIChecker struct {

	// respone length of a known bad route
	BadLength int

	// this is the HTTP client that will be conducting the
	// requests to the target.
	Checker LFIClient

	// response length of a known good route.
	GoodLength int

	// a route that will return a "200 OK" response. this route
	// will be used in various locations.
	GoodRoute string

	// a route that will return a "404 Not Found" response. this route
	// will be used in various locations to check for LFI.
	BadRoute string

	// LFI options associated with this checker
	Options LFIOptions
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
	Parameters []string

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

