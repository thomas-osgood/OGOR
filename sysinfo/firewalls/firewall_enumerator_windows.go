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
		activefirewalls = append(activefirewalls, "Windows Defender")
	}

	return activefirewalls, nil
}
