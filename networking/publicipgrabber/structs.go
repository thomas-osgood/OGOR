package publicipgrabber

import "net/http"

// structure representing the response from api.whatismyip.com/app.php
// when querying information related to an IP address.
type AppResponse struct {
	// data returned from ip2location.com query made
	// by api.whatismyip.com/app.php.
	Ip2location LocationResponse `json:"ip2location.com"`

	// data returned from ipdata.co query made
	// by api.whatismyip.com/app.php.
	Ipdata LocationResponse `json:"ipdata.co"`
}

// structure representing the response from api.whatismyip.com/app.php when
// querying DNS information related to a URL.
type DnsResponse struct {
	Arecords []string `json:"a-records"`
}

// structure representing an error response from api.whatismyip.com.
type ErrorResponse struct {
	Error string `json:"error"`
}

// structure defining the IP2Location and IPData
// JSON returns from app.php.
type LocationResponse struct {
	Asn        string `json:"asn"`
	City       string `json:"city"`
	Region     string `json:"region"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
	Isp        string `json:"isp"`
	TimeZone   string `json:"time_zone"`
}

// structure defining a PublicIPGrabber object. this
// will have associated functions to query the site
// api.whatismyip.com and grab the public IP info.
type PublicIPGrabber struct {
	// http client that will be used to carry
	// out queries to api.whatismyip.com.
	client *http.Client

	// information pulled down from api.whatismyip.com.
	PublicIP PublicIPInfo
}

// structure defining the object that will be used
// to initialize a public ip grabber object.
type PublicIPGrabberOptions struct {
	// http client that will be used to carry
	// out queries to api.whatismyip.com.
	Client *http.Client
}

// structure holding public ip information. this will
// be used in the PublicIPGrabber object and associated
// request to api.whatismyip.com.
type PublicIPInfo struct {
	// public ip address
	Ip string `json:"ip"`

	// geolocation of the server hosting the ip
	Location string `json:"geo"`

	// provider hosting the IP address
	Provider string `json:"isp"`
}
