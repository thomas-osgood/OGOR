//go:build windows
// +build windows

package firewalls

// this string slice represents the
// arguments that will be passed to
// the command to enumerate the
// firewall binaries.
var enumarguments []string = []string{"/R", "C:\\", ""}
