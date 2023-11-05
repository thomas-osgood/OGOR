package firewalls

import "github.com/thomas-osgood/OGOR/output"

// enum representing the possible statuses of
// a firewall (disabled, enabled, unknown).
type enumBase int

// structure representing the object that will
// conduct the firewall enumeration.
type FirewallEnumerator struct {
	firewalls map[string]enumBase
	formatter *output.Formatter
	printer   *output.Outputter
	services  map[string]bool
}
