package firewalls

import "github.com/thomas-osgood/OGOR/output"

// function designed to create and initialize a new
// FirewallEnumerator object that can be used by the
// end-user. this returns a pointer to the newly
// created object and an error (nil if successful).
func NewFirewallEnumerator(optFuncs ...FirewallEnumOptFunc) (enumerator *FirewallEnumerator, err error) {
	var defaultOptions FirewallEnumOptions = FirewallEnumOptions{DisplayErrors: true}
	var fn FirewallEnumOptFunc

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

	// go through the options set by the user and set
	// the variables of the new FirewallEnumerator based
	// on the user's input(s).
	for _, fn = range optFuncs {
		err = fn(&defaultOptions)
		if err != nil {
			return nil, err
		}
	}

	// set the new enumerator variable values based on
	// the user input.
	enumerator.displayErrors = defaultOptions.DisplayErrors

	return enumerator, nil
}

// function designed to turn off the error output for
// a FirewallEnumerator. this will not show non-critical
// errors when the occur; they will just be ignored and
// the function will continue without signaling to the
// user that something went wrong.
func ErrorOutputOff() FirewallEnumOptFunc {
	return func(opts *FirewallEnumOptions) error {
		opts.DisplayErrors = false
		return nil
	}
}
