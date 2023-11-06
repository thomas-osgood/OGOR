//go:build !windows
// +build !windows

package firewalls

import (
	"context"
	"os/exec"
	"strings"
)

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
	var currentBinary string
	var active bool

	if (fe.firewalls == nil) || (len(fe.firewalls) < 1) {
		err = fe.EnumBinaries()
		if err != nil {
			return nil, err
		}
	}

	// go through each discovered firewall binary
	// and test to see if the firewall is active.
	for currentBinary = range fe.firewalls {
		switch strings.ToLower(currentBinary) {
		case "ufw":
			active, err = fe.checkUFW()
			if err != nil {
				continue
			}
		case "fiewall-cmd":
			active, err = fe.checkFirewallCmd()
			if err != nil {
				continue
			}
		default:
			continue
		}

		if active {
			fe.firewalls[currentBinary] = Enabled
		} else {
			fe.firewalls[currentBinary] = Disabled
		}
	}

	return activefirewalls, nil
}

// function designed to determine whether the
// Uncomplicated Firewall (UFW) is currently
// active on the machine.
//
// this will return a boolean indication of
// active, along with an error. if everything
// executes as expected, the error will be nil.
func (fe *FirewallEnumerator) checkUFW() (active bool, err error) {
	var cancel context.CancelFunc
	var cmd *exec.Cmd
	var cmdctx context.Context
	const command string = "ufw"
	var commandargs []string = []string{"status"}
	var outbytes []byte
	var outstring string

	cmdctx, cancel = context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	cmd = exec.CommandContext(cmdctx, command, commandargs...)
	outbytes, err = cmd.Output()
	if err != nil {
		return false, err
	}

	outstring = string(outbytes)
	outstring = strings.TrimSpace(outstring)

	return active, nil
}

// function designed to determine whether the
// firewall-cmd output indicates that any firewalls
// are currently active.
func (fe *FirewallEnumerator) checkFirewallCmd() (active bool, err error) {
	var cancel context.CancelFunc
	var cmd *exec.Cmd
	var cmdctx context.Context
	const command string = "firewall-cmd"
	var commandargs []string = []string{"--state"}
	var outbytes []byte
	var outstring string

	cmdctx, cancel = context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	cmd = exec.CommandContext(cmdctx, command, commandargs...)
	outbytes, err = cmd.Output()
	if err != nil {
		return false, err
	}

	outstring = string(outbytes)
	outstring = strings.TrimSpace(outstring)

	// read the "firewall-cmd --state" output and
	// set the active flag.
	switch strings.ToLower(outstring) {
	case "not running":
		active = false
	case "running":
		active = true
	default:
		active = false
	}

	return active, nil
}

// function designed to grab all running services
// the user can see and save the service name and
// status in the services map.
func (fe *FirewallEnumerator) getServices() (err error) {
	var cancel context.CancelFunc
	var cmd *exec.Cmd
	var cmdctx context.Context
	const command string = "service"
	var commandargs []string = []string{"--status-all"}
	var currentline string
	var indicator string
	var outbytes []byte
	var outstring string
	var servicename string
	var servicestatus bool
	var splitline []string
	var splitstring []string

	cmdctx, cancel = context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	cmd = exec.CommandContext(cmdctx, command, commandargs...)
	outbytes, err = cmd.Output()
	if err != nil {
		return err
	}

	outstring = string(outbytes)
	outstring = strings.TrimSpace(outstring)

	splitstring = strings.Split(outstring, "\n")

	// go through each line and determine the service
	// name and status.
	for _, currentline = range splitstring {
		splitline = strings.Split(currentline, "]")
		indicator = strings.TrimSpace(strings.Replace(splitline[0], "[", "", 1))
		servicename = strings.TrimSpace(splitline[1])

		// determine the service status based on the
		// first piece of the line ([ + ] or [ - ]).
		switch indicator {
		case "+":
			servicestatus = true
		default:
			servicestatus = false
		}

		fe.services[servicename] = servicestatus
	}

	return nil
}
