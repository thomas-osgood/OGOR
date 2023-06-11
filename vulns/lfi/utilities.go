package lfichecker

import "net/http"

// function designed to create and initialize a new LFI/Directory Traversal
// checker and return a pointer to it to the user. this returns a pointer
// to the LFIChecker object and nil if no error occurs, otherwise it returns
// nil and an error.
func NewLFIChecker(baseurl string, usropts ...LFIOptsFunc) (checker *LFIChecker, err error) {
	var client http.Client = http.Client{}
	var opts LFIOptions = LFIOptions{Parameters: make(map[string]string), DoubleEncoding: false, SSLConnection: false}

	// set any options passed in by the user.
	for _, fnc := range usropts {
		fnc(&opts)
	}

	// create the new LFIChecker object to return to the user.
	checker = &LFIChecker{
		Checker:          LFIClient{baseurl: baseurl, client: client},
		GoodRoute:        "",
		BadRoute:         "",
		BadLengthParams:  make(map[string]int),
		BlankLength:      make(map[string]int),
		Options:          opts,
		VulnerableParams: make(map[string]string),
	}

	return checker, nil
}
