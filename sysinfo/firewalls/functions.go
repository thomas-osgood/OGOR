package firewalls

import "github.com/thomas-osgood/OGOR/output"

// function designed to create and initialize a new
// FirewallEnumerator object that can be used by the
// end-user. this returns a pointer to the newly
// created object and an error (nil if successful).
func NewFirewallEnumerator() (enumerator *FirewallEnumerator, err error) {
	enumerator = new(FirewallEnumerator)

	// create the outputter object. this will be used
	// to nicely print information.
	enumerator.printer, err = output.NewOutputter()
	if err != nil {
		return nil, err
	}

	// create the formatter object. this will be used
	// to nicely print information.
	enumerator.formatter, err = output.NewFormatter()
	if err != nil {
		return nil, err
	}

	// initialize the firewalls map so a nil error
	// does not get thrown when attempting to assign
	// to, or read from, it.
	enumerator.firewalls = make(map[string]enumBase)

	// initialize the services map so a nil error does
	// not get thrown when attempting to assign to, or
	// read from, it.
	enumerator.services = make(map[string]bool)

	return enumerator, nil
}
