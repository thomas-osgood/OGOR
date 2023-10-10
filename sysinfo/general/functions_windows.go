//go:build windows
// +build windows

package general

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// function designed to enumerate mounts on a windows
// machine.
func EnumMounts() (mounts []MountInfo, err error) {
	mounts = make([]MountInfo, 0)
	return mounts, nil
}

// function deisgned to read the cpuinfo file and pull
// out information of interest. this will help give a
// better understanding of the machine's CPU(s).
//
// for ref:
// https://winaero.com/get-cpu-information-via-command-prompt-in-windows-10/
func GetCPUInfo() (info AllCpuInfo, err error) {
	const target string = "wmic"

	var cmd *exec.Cmd
	var cmdstr string = "cmd.exe"
	var cmdarg []string = []string{"/C", target, "cpu", "get", "caption,deviceid,name,numberofcores,maxclockspeed,status", "/format:csv"}
	var linesplit []string
	var outbytes []byte
	var outsplit [][]byte
	var speed float64

	// initialize Cpus slice in AllCpuInfo struct
	info.Cpus = make([]CpuInfo, 0)

	cmd = exec.Command(cmdstr, cmdarg...)
	outbytes, err = cmd.CombinedOutput()
	if err != nil {
		return AllCpuInfo{}, err
	}

	outsplit = bytes.Split(bytes.TrimSpace(outbytes), []byte("\n"))
	if len(outsplit) < 2 {
		return AllCpuInfo{}, fmt.Errorf("no cpu information discovered")
	}

	for _, outbytes = range outsplit[1:] {
		linesplit = strings.Split(string(bytes.TrimSpace(outbytes)), ",")
		fmt.Printf("[%+v]\n", linesplit)
		speed, err = strconv.ParseFloat(linesplit[3], 64)
		if err != nil {
			speed = -1
		}
		info.Cpus = append(info.Cpus, CpuInfo{ProcessorModel: linesplit[4], ProcessorSpeed: speed, ProcessorVendor: linesplit[1]})
	}

	return info, nil
}
