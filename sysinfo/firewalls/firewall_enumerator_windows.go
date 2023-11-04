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

// function designed to look for the target
// firewall binaries on the current machine.
//
// this will return a string slice outlining
// the absolute paths of any binaries found.
func (fe *FirewallEnumerator) EnumBinaries() (err error) {
	var cancel context.CancelFunc
	var cmd *exec.Cmd
	var cmdctx context.Context
	const cmdtimeout time.Duration = 10 * time.Second
	var currentbin string
	const enumcommand string = "where"
	var enumarguments []string = make([]string, 1)
	var outbytes []byte
	var outstr string

	// get the most up-to-date list of services.
	err = fe.getServices()
	if err != nil {
		return err
	}

	// loop through each target binary and run "which <targetbin>".
	// if this command does not error out and has a length greater
	// than zero after the space has been trimmed, it will be appended
	// to the binaries slice.
	for _, currentbin = range targetbinaries {
		enumarguments[0] = currentbin

		cmdctx, cancel = context.WithTimeout(context.Background(), cmdtimeout)
		defer cancel()

		// execute the "which <targetbin>" command with a 10 second
		// timeout. if the command does not return in this time (or
		// if it errors out) the loop will move onto the next binary.
		cmd = exec.CommandContext(cmdctx, enumcommand, enumarguments...)
		outbytes, err = cmd.Output()
		if err != nil {
			continue
		}

		// remove the leading and trailing newlines and spaces, then
		// check the string length. if hte string length is less than
		// one, continue on with the loop because there was no binary
		// discovered by the command and a blank line was returned.
		outstr = strings.TrimSpace(string(outbytes))
		if len(outstr) < 1 {
			continue
		}

		fe.firewalls[outstr] = Unknown

	}

	// if the length of the return slice is zero, raise an error
	// saying no firewall binaries were discovered.
	if len(fe.firewalls) < 1 {
		return errors.New("no firewall binaries discovered")
	}

	return nil
}
