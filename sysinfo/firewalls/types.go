package firewalls

// enum representing the possible statuses of
// a firewall (disabled, enabled, unknown).
type enumBase int

// structure representing the object that will
// conduct the firewall enumeration.
type FirewallEnumerator struct {
	firewalls map[string]enumBase
	services  map[string]bool
}
