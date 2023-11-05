package firewalls

import "time"

// default timeout used in commands.
const DefaultTimeout time.Duration = 10 * time.Second

// enum defining the states a fiewall can be in.
const (
	Disabled enumBase = 0
	Enabled  enumBase = 1
	Unknown  enumBase = 2
)
