package firewalls

import "github.com/thomas-osgood/OGOR/output"

// enum representing the possible statuses of
// a firewall (disabled, enabled, unknown).
type enumBase int

// structure representing the object that will
// conduct the firewall enumeration.
type FirewallEnumerator struct {
	displayErrors bool
	firewalls     map[string]enumBase
	formatter     *output.Formatter
	printer       *output.Outputter
	services      map[string]bool
}

// structure designed to hold the options a user
// can set when initializing a FirewallEnumerator.
type FirewallEnumOptions struct {
	DisplayErrors bool
}

// type definition explaining a Firewall Enumerator
// Options Function. this will be used during the
// initialization of a new FirewalEnumerator to set
// user-controlled variables.
type FirewallEnumOptFunc func(*FirewallEnumOptions) error
