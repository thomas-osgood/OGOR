package firewalls

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

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
	var outbytes []byte
	var outstr string

	// get the most up-to-date list of services.
	//
	// if an error occurs, this will not be considered "critical"
	// and will not cause a return; instead, a message will be
	// output if "displayErrors" is set to true.
	err = fe.getServices()
	if err != nil {
		if fe.displayErrors {
			fe.printer.ErrMsg(err.Error())
		}
	}

	// loop through each target binary and run "which <targetbin>".
	// if this command does not error out and has a length greater
	// than zero after the space has been trimmed, it will be appended
	// to the binaries slice.
	for _, currentbin = range targetbinaries {
		enumarguments[len(enumarguments)-1] = currentbin

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
