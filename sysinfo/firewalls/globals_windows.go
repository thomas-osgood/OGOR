//go:build windows
// +build windows

package firewalls

// this string slice represents the
// arguments that will be passed to
// the command to enumerate the
// firewall binaries.
//
// ref: https://ss64.com/nt/where.html
var enumarguments []string = []string{"/R", "C:\\", ""}
