package ipgrabber

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

// function designed to add an IP address to the list of
// discovred IP addresses. this first checks if the address
// was previously discovered/added to the list. if the
// address is already in the list, it does not add it and
// returns an error.
func (grabber *IPGrabber) AddIP(ipaddr net.IP) (err error) {
	if err = grabber.AlreadyFound(ipaddr); err == nil {
		return errors.New("address already discovered. not adding.")
	}

	grabber.Discovered = append(grabber.Discovered, ipaddr)

	return nil
}

// function designed to determine if the given IP address is already
// in the list of discovred addresses. returns error if the address
// is not in the list, nil if it is in the list.
func (grabber *IPGrabber) AlreadyFound(ipaddr net.IP) (err error) {
	var found bool = false

	for _, curip := range grabber.Discovered {
		if ipaddr.String() == curip.String() {
			found = true
		}
	}

	if !found {
		return errors.New("address not previously discovred")
	}

	return nil
}

// function designed to check wheter an interface with a given name
// exists in the discovered interfaces.
func (grabber *IPGrabber) InterfaceExists(ifaceName string) (err error) {
	var found bool = false

	for _, iface := range grabber.Interfaces {
		if iface.Name == ifaceName {
			found = true
			break
		}
	}

	if !found {
		return errors.New(fmt.Sprintf("no interface named \"%s\" found in discovered interfaces", ifaceName))
	}

	return nil
}

// function designed to populate/repopulate the Interfaces slice
// of the IPGrabber struct by calling the net.Interfaces function.
func (grabber *IPGrabber) GrabInterfaces() (err error) {
	//------------------------------------------------------------
	// grab all network interfaces
	//------------------------------------------------------------
	grabber.Interfaces, err = net.Interfaces()
	if err != nil {
		return err
	}

	return nil
}

// function designed to acquire the IP address of a target interface.
// this will return the net.IP object and nil if the interface is
// successfully discovered, otherwise it will return nil and error.
func (grabber *IPGrabber) GrabInterfaceIP(targetiface string) (ipaddr net.IP, err error) {
	var addresses []net.Addr
	var address net.Addr
	var foundiface bool = false

	if grabber.IFacesEmpty() {
		err = grabber.GrabInterfaces()
		if err != nil {
			return nil, err
		}
	}

	for _, iface := range grabber.Interfaces {
		//------------------------------------------------------------
		// ignore loopback addresses (127.0.0.1)
		//------------------------------------------------------------
		if strings.Contains(iface.Flags.String(), net.FlagLoopback.String()) {
			continue
		}

		//------------------------------------------------------------
		// if target interface is set, only display target
		//------------------------------------------------------------
		if (len(targetiface) > 0) && (iface.Name != targetiface) {
			continue
		}

		foundiface = true

		//------------------------------------------------------------
		// grab all addresses from current interface
		//------------------------------------------------------------
		addresses, err = iface.Addrs()
		if err != nil {
			return nil, err
		}

		//------------------------------------------------------------
		// loop through all addresses present in current interface
		//------------------------------------------------------------
		for _, address = range addresses {
			switch v := address.(type) {
			case *net.IPNet:
				ipaddr = v.IP
			case *net.IPAddr:
				ipaddr = v.IP
			}

			//------------------------------------------------------------
			// only grab IPv4 addresses
			//------------------------------------------------------------
			if ipaddr.To4() == nil {
				continue
			}

			break
		}

		// interface found: break loop and return
		break
	}

	//------------------------------------------------------------
	// error finding target (or any) network interface
	//------------------------------------------------------------
	if (len(targetiface) > 0) && !foundiface {
		return nil, errors.New(fmt.Sprintf("unable to find interface \"%s\"", targetiface))
	} else if !foundiface {
		return nil, errors.New("no network interfaces discovered")
	}

	return ipaddr, nil
}

//    Function designed to acquire all IPv4 network addresses
//    attached to the current machine.
func (grabber *IPGrabber) GrabIPs() (err error) {
	var address net.Addr
	var addresses []net.Addr
	var iface net.Interface
	var ip net.IP

	if grabber.IFacesEmpty() {
		err = grabber.GrabInterfaces()
		if err != nil {
			return err
		}
	}

	for _, iface = range grabber.Interfaces {

		//------------------------------------------------------------
		// ignore loopback addresses (127.0.0.1)
		//------------------------------------------------------------
		if strings.Contains(iface.Flags.String(), net.FlagLoopback.String()) {
			continue
		}

		//------------------------------------------------------------
		// grab all addresses from current interface
		//------------------------------------------------------------
		addresses, err = iface.Addrs()
		if err != nil {
			return err
		}

		//------------------------------------------------------------
		// loop through all addresses present in current interface
		//------------------------------------------------------------
		for _, address = range addresses {
			switch v := address.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			//------------------------------------------------------------
			// only grab IPv4 addresses
			//------------------------------------------------------------
			if ip.To4() == nil {
				continue
			}

			// add IP address to list of discovred IPs.
			//
			// ignore error return from this, because all it means
			// is that the IP address already exists in the list
			// and was not added.
			grabber.AddIP(ip)
		}
	}

	return nil
}

// function designed to check whether the interface list is empty.
// returns true/false based on the length of grabber.Interfaces.
func (grabber *IPGrabber) IFacesEmpty() (empty bool) {
	if len(grabber.Interfaces) < 1 {
		empty = true
	} else {
		empty = false
	}
	return empty
}

//
// Function Name: ListIPs
//
// Author: Thomas Osgood
//
// Description:
//
//    Function designed to grab all IPv4 IPs attached to all
//    network interfaces and list them out.
//
// Input(s):
//
//    None
//
// Return(s):
//
//    err - error. error or nil.
//
func (grabber *IPGrabber) ListIPs() (err error) {

	err = grabber.GrabIPs()
	if err != nil {
		return err
	}

	for _, ipAddr := range grabber.Discovered {
		fmt.Printf("%s\n", ipAddr.String())
	}

	return nil
}

