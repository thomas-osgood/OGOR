package firewalls

import "time"

// constant array holds the common binaries
// used for firewall administration.
var targetbinaries = [...]string{"firewall-cmd", "iptables", "ufw"}

// default timeout used in commands.
const DefaultTimeout time.Duration = 10 * time.Second

// enum defining the states a fiewall can be in.
const (
	Disabled enumBase = 0
	Enabled  enumBase = 1
	Unknown  enumBase = 2
)
