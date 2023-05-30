package ipgrabber

import "net"

// structure representing the main object that will search for
// and hold the discovered addresses and interfaces.
type IPGrabber struct {
	Interfaces []net.Interface
	Discovered []net.IP
}

