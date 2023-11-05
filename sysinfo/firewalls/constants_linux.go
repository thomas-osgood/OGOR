//go:build !windows
// +build !windows

package firewalls

// constant array holds the common binaries
// used for firewall administration.
var targetbinaries = [...]string{"firewall-cmd", "iptables", "ufw"}
