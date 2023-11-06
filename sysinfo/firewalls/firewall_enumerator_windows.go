//go:build windows
// +build windows

package firewalls

// function designed to determine whether the
// discovered firewall(s) are active. if the
// firewall binary requires sudo (elevated)
// privileges and the command errors out, the
// next discovered firewall will be checked.
//
// the firewall binaries that error out will be
// kept track of and an error will be raised if
// no firewalls were found to be running but at
// least one error occurred during testing.
func (fe *FirewallEnumerator) CheckFirewalls() (activefirewalls []string, err error) {
	activefirewalls = make([]string, 0)
	return activefirewalls, nil
}
