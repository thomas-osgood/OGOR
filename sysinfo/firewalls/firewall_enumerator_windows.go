//go:build windows
// +build windows

package firewalls

// function designed to determine whether the
// Uncomplicated Firewall (UFW) is currently
// active on the machine.
//
// this will return a boolean indication of
// active, along with an error. if everything
// executes as expected, the error will be nil.
func (fe *FirewallEnumerator) checkUFW() (active bool, err error) {
	return false, nil
}

// function designed to grab all running services
// the user can see and save the service name and
// status in the services map.
func (fe *FirewallEnumerator) getServices() (err error) {
	return nil
}
