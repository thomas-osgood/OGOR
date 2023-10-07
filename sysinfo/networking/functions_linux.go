//go:build !windows
// +build !windows

package networking

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// function designed to read the resolv.conf file
// and return the IP addresses of the dns nameservers.
func EnumDnsServers() (servers []string, err error) {
	var cmd *exec.Cmd
	var cmdstr string = "cat"
	var cmdarg []string = []string{"/etc/resolv.conf"}
	var line string
	var linesplit []string
	var outbytes []byte
	var outsplit [][]byte

	servers = make([]string, 0)

	cmd = exec.Command(cmdstr, cmdarg...)

	outbytes, err = cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	outbytes = bytes.TrimSpace(outbytes)
	if len(outbytes) < 1 {
		return nil, fmt.Errorf("nothing found in /etc/resolv.conf")
	}

	outsplit = bytes.Split(outbytes, []byte("\n"))

	// loop through each line of the resolv.conf
	// file and pull out the information that is
	// relevant to this function.
	for _, outbytes = range outsplit {
		line = string(bytes.TrimSpace(outbytes))
		linesplit = strings.Split(line, " ")

		if len(linesplit) < 2 {
			continue
		}

		if linesplit[0] != "nameserver" {
			continue
		}

		servers = append(servers, linesplit[1])
	}

	if len(servers) < 1 {
		return nil, fmt.Errorf("nothing found in /etc/resolv.conf")
	}

	return servers, nil
}
