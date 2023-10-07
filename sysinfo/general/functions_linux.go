package general

import (
	"bytes"
	"os/exec"
	"strings"
)

// function designed to discover the mounts on a
// linux system. this will give the name, location,
// type, and privileges.
//
// this uses the "mount" command to list all the mounts
// the user can see, then parses the information and
// returns a slice of MountInfo structs.
func EnumMounts() (mounts []MountInfo, err error) {
	var cmd *exec.Cmd
	var cmdstr string = "mount"
	var cmdarg []string = []string{}
	var line string
	var linesplit []string
	var outbytes []byte
	var outsplit [][]byte

	mounts = make([]MountInfo, 0)

	cmd = exec.Command(cmdstr, cmdarg...)

	outbytes, err = cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	outsplit = bytes.Split(bytes.TrimSpace(outbytes), []byte("\n"))

	for _, outbytes = range outsplit {
		line = string(bytes.TrimSpace(outbytes))
		linesplit = strings.Split(line, " ")

		mounts = append(
			mounts,
			MountInfo{
				Name:       linesplit[0],
				Location:   linesplit[2],
				Type:       linesplit[4],
				Privileges: linesplit[5],
			},
		)
	}

	return mounts, nil
}
