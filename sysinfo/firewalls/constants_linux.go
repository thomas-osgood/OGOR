//go:build !windows
// +build !windows

package firewalls

// command that will be used to enumerate
// the firewall binaries.
const enumcommand string = "which"

// constant array holds the common binaries
// used for firewall administration.
var targetbinaries = [...]string{"firewall-cmd", "iptables", "ufw"}
