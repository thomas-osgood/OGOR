//go:build windows
// +build windows

package firewalls

import (
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// function designed to determine whether the
// firewall enumeration output indicates that
// any firewalls are currently active.
func (fe *FirewallEnumerator) checkFirewallState() (active bool, err error) {
	var cancel context.CancelFunc
	var cmd *exec.Cmd
	var cmdctx context.Context
	const command string = "netsh"
	var commandargs []string = []string{"advfirewall", "show", "currentprofile"}
	var outbytes []byte
	var outstring string
	var regpat string = "state\\s+.*\\n"
	var re *regexp.Regexp
	var rematches []string

	cmdctx, cancel = context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	cmd = exec.CommandContext(cmdctx, command, commandargs...)
	outbytes, err = cmd.Output()
	if err != nil {
		return false, err
	}

	outstring = string(outbytes)
	outstring = strings.TrimSpace(outstring)

	re, err = regexp.Compile(regpat)
	if err != nil {
		return false, err
	}

	rematches = re.FindAllString(strings.ToLower(outstring), -1)
	if rematches == nil {
		return false, fmt.Errorf("no firewall state output detected")
	}

	// read the "netsh advfirewall show currentprofile"
	// output and set the active flag.
	switch strings.ToLower(outstring) {
	case "off":
		active = false
	default:
		active = true
	}

	return active, nil
}

// function designed to get all the running services on
// the current machine.
func (fe *FirewallEnumerator) getServices() (err error) {
	var cancel context.CancelFunc
	var checkcmd *exec.Cmd
	var clsplit []string
	var cmd *exec.Cmd
	var cmdctx context.Context
	const command string = "sc"
	var commandargs []string = []string{"query"}
	var currentline string
	var outbytes []byte
	var outstring string
	var outstringlines []string
	var splitstate []string
	var statematches []string
	const statepat string = "STATE\\s+:\\s[0-9]\\s+[a-zA-Z]+"
	var statereg *regexp.Regexp
	var state string
	const svcpat string = "SERVICE_NAME:\\s+.*"
	var svcreg *regexp.Regexp

	statereg, err = regexp.Compile(statepat)
	if err != nil {
		return err
	}

	svcreg, err = regexp.Compile(svcpat)
	if err != nil {
		return err
	}

	cmdctx, cancel = context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	cmd = exec.CommandContext(cmdctx, command, commandargs...)
	outbytes, err = cmd.Output()
	if err != nil {
		return err
	}

	outstring = string(outbytes)
	outstring = strings.TrimSpace(outstring)

	outstringlines = svcreg.FindAllString(outstring, -1)
	if outstringlines == nil {
		return fmt.Errorf("no services discovered")
	}

	for _, currentline = range outstringlines {
		currentline = strings.TrimSpace(currentline)
		if len(currentline) == 0 {
			continue
		}
		clsplit = strings.Split(strings.Split(currentline, ":")[1], " ")
		currentline = strings.TrimSpace(clsplit[len(clsplit)-1])

		checkcmd = exec.Command("sc", "query", currentline)
		outbytes, err = checkcmd.Output()
		if err != nil {
			fe.services[currentline] = false
			continue
		}

		statematches = statereg.FindAllString(string(outbytes), -1)
		if statematches == nil {
			fe.services[currentline] = false
			continue
		}

		// list of states to check for in future:
		// https://learn.microsoft.com/en-us/windows/win32/api/winsvc/ns-winsvc-service_status
		state = strings.TrimSpace(strings.Split(statematches[0], ":")[1])
		splitstate = strings.Split(state, " ")
		state = strings.TrimSpace(splitstate[len(splitstate)-1])
		if state == "RUNNING" {
			fe.services[currentline] = true
		} else {
			fe.services[currentline] = false
		}
	}

	return nil
}

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
	var active bool

	activefirewalls = make([]string, 0)

	active, err = fe.checkFirewallState()
	if err != nil {
		return nil, err
	}

	if active {
		activefirewalls = append(activefirewalls, "Windows Defender Firewall")
	}

	return activefirewalls, nil
}
