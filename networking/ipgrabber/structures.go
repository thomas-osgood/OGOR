package ipgrabber

import "net"

// structure representing the main object that will search for
// and hold the discovered addresses and interfaces.
type IPGrabber struct {

	// slice holding all the discovered network interfaces on
	// the current system. this excludes the loopback interface.
	Interfaces []net.Interface

	// slice holding all the discovered IP addresses for all
	// the interfaces in the Interfaces slice.
	Discovered []net.IP
}

