package fingerprinter

import (
	"errors"
	"net/http"
	"strings"
)

// function designed to make a request to the target site and attempt
// to pull the ALLOW header information from it. if the ALLOW header
// is present, the data will be parsed and the allowedmethods slice
// will be populated.
func (f *Fingerprinter) GetAllowedMethods() (err error) {
	var client http.Client = http.Client{}
	var methodheader string
	var resp *http.Response
	var splitheader []string

	// clear any existing values in allowedmethods slice or
	// init a new string slice.
	f.allowedmethods = []string{}

	resp, err = client.Head(f.Target)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	methodheader = resp.Header.Get("Allowed")
	if len(methodheader) < 1 {
		return errors.New("Allowed header not present in HEAD request. cannot determine allowed methods.")
	}
	methodheader = strings.ToUpper(strings.ReplaceAll(methodheader, " ", ""))

	splitheader = strings.Split(methodheader, ",")

	for _, method := range splitheader {
		switch {
		case method == http.MethodHead:
			f.allowedmethods = append(f.allowedmethods, http.MethodHead)
		case method == http.MethodGet:
			f.allowedmethods = append(f.allowedmethods, http.MethodGet)
		case method == http.MethodPost:
			f.allowedmethods = append(f.allowedmethods, http.MethodPost)
		case method == http.MethodPut:
			f.allowedmethods = append(f.allowedmethods, http.MethodPut)
		case method == http.MethodPatch:
			f.allowedmethods = append(f.allowedmethods, http.MethodPatch)
		case method == http.MethodOptions:
			f.allowedmethods = append(f.allowedmethods, http.MethodOptions)
		case method == http.MethodDelete:
			f.allowedmethods = append(f.allowedmethods, http.MethodDelete)
		case method == http.MethodPatch:
			f.allowedmethods = append(f.allowedmethods, http.MethodPatch)
		case method == http.MethodTrace:
			f.allowedmethods = append(f.allowedmethods, http.MethodTrace)
		case method == http.MethodConnect:
			f.allowedmethods = append(f.allowedmethods, http.MethodConnect)
		}
	}

	return nil
}

// function designed to acquire the type of server hosting the target
// by making a request to the target and analyzing the headers that
// get returned.
func (f *Fingerprinter) GetServerType() (err error) {
	var client http.Client = http.Client{}
	var resp *http.Response
	var serverheader string

	resp, err = client.Get(f.Target)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	serverheader = strings.Trim(resp.Header.Get("server"), " ")
	if len(serverheader) < 1 {
		serverheader = "unknown"
	}

	f.servertype = serverheader

	return nil
}
