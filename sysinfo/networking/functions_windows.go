//go:build windows
// +build windows

package networking

// function designed to enumerate the DNS information
// of a Windows machine and return the nameserver IPs.
func EnumDnsServers() (servers []string, err error) {
	servers = make([]string, 0)
	return servers, nil
}
