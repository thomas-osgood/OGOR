//go:build windows
// +build windows

package firewalls

// command that will be used to enumerate
// the firewall binaries.
const enumcommand string = "where"

// constant array holds the common binaries
// used for firewall administration.
var targetbinaries = [...]string{"mpcmdrun.exe"}
