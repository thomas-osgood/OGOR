package firewalls

import (
	"errors"
	"fmt"
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

// function designed to output the discovered firewall binaries.
//
// if no binaries have been discovered, an error will be returned.
func (fe *FirewallEnumerator) ShowFirewallBinaries() (err error) {
	var currentbinary string

	if (fe.firewalls == nil) || (len(fe.firewalls) < 1) {
		return errors.New("no binaries discovered. try running EnumBinaries")
	}

	for currentbinary = range fe.firewalls {
		fe.printer.SucMsg(fmt.Sprintf("Binary: %s", currentbinary))
	}

	return nil
}

// function designed to output the discovered services and their status.
//
// if no services have been discovered an error will be returned.
func (fe *FirewallEnumerator) ShowServices() (err error) {
	var currentservice string
	var currentstatus bool

	if (fe.services == nil) || (len(fe.services) < 1) {
		return errors.New("no services discovered")
	}

	fe.printer.InfMsg(fmt.Sprintf("%-35s|%-10s", "service", "active"))
	fe.printer.InfMsg(fe.formatter.PrintChar('-', 46))
	for currentservice, currentstatus = range fe.services {
		fe.printer.InfMsg(fmt.Sprintf("%-35s|%t", currentservice, currentstatus))
	}

	return nil
}
