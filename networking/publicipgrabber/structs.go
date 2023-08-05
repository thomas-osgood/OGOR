package publicipgrabber

// structure defining a PublicIPGrabber object. this
// will have associated functions to query the site
// api.whatismyip.com and grab the public IP info.
type PublicIPGrabber struct {
	PublicIP PublicIPInfo
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
