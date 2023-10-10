//go:build windows
// +build windows

package networking

import (
	"bytes"
	"os/exec"
	"strings"
)

// function designed to read the resolv.conf file
// and return the IP addresses of the dns nameservers.
//
// for ref:
// https://kevincurran.org/com320/labs/dns.htm
func EnumDnsServers() (servers []string, err error) {
	var cmd *exec.Cmd
	var cmdstr string = "cmd.exe"
	var cmdarg []string = []string{"/C", "ipconfig", "/all"}
	var dnsip string
	var dnsips string
	var outbytes []byte
	var outsplit [][]byte

	servers = make([]string, 0)

	cmd = exec.Command(cmdstr, cmdarg...)
	outbytes, err = cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	outbytes = bytes.TrimSpace(outbytes)

	outsplit = bytes.Split(outbytes, []byte("\n"))
	for _, outbytes = range outsplit {
		if bytes.Contains(outbytes, []byte("DNS Servers")) {
			dnsips = string(bytes.Join(bytes.Split(outbytes, []byte(":"))[1:], []byte(",")))
			for _, dnsip = range strings.Split(dnsips, ",") {
				servers = append(servers, dnsip)
			}
		}
	}

	return servers, nil
}
