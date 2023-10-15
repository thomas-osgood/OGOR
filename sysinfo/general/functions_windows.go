//go:build windows
// +build windows

package general

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// function designed to enumerate mounts on a windows
// machine.
func EnumMounts() (mounts []MountInfo, err error) {
	mounts = make([]MountInfo, 0)
	return mounts, nil
}

// function designed to get the system architecture (32-bit or 64-bit)
func GetArchitecture() (architecture string, err error) {
	var cmd *exec.Cmd
	var cmdstr string = "wmic"
	var cmdarg []string = []string{"OS", "get", "OSArchitecture", "/VALUE"}
	var outbytes []byte

	cmd = exec.Command(cmdstr, cmdarg...)
	outbytes, err = cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	outbytes = bytes.TrimSpace(outbytes)

	if len(outbytes) < 1 {
		return "", fmt.Errorf("no output. unable to determine architecture")
	}

	architecture = strings.Split(string(outbytes), "=")[1]

	return architecture, nil
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

// function designed to get the kernel version.
func GetKernelVersion() (version string, err error) {
	var cmd *exec.Cmd
	var cmdstr string = "systeminfo"
	var matches [][]byte
	var outbytes []byte
	var re *regexp.Regexp
	const searchpat string = `OS\sVersion:.*\n`

	cmd = exec.Command(cmdstr)
	outbytes, err = cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	outbytes = bytes.TrimSpace(outbytes)

	if len(outbytes) < 1 {
		return "", fmt.Errorf("no output. unable to determine kernel version")
	}

	re, err = regexp.Compile(searchpat)
	if err != nil {
		return "", err
	}

	matches = re.FindAll(outbytes, -1)
	if matches == nil {
		return "", fmt.Errorf("unable to find OS Version in SystemInfo")
	}

	version = strings.Split(strings.TrimSpace(strings.Split(string(matches[0]), ":")[1]), " ")[0]

	return version, nil
}

// function designed to grab the general system information
// for a windows machine.
func GetSysInfo() (info BasicSysInfo, err error) {
	var cmd *exec.Cmd
	var cmdstr string = "systeminfo"
	var cmdarg []string = []string{"/fo", "csv", "/nh"}
	var curline string
	var hotfix string
	var hotfixes []string
	var i int
	var outbytes []byte
	var outsplit []string

	info.OperatingSystem = OSInfo{}
	info.OperatingSystem.Hotfixes = make([]string, 0)

	cmd = exec.Command(cmdstr, cmdarg...)
	outbytes, err = cmd.CombinedOutput()
	if err != nil {
		return BasicSysInfo{}, err
	}

	outsplit = strings.Split(string(outbytes), "\",")
	for i = range outsplit {
		outsplit[i] = strings.TrimLeft(outsplit[i], "\"")
	}

	// if there are hotfixes present, loop through them
	// and add them to the Hotfixes slice.
	hotfixes = strings.Split(outsplit[30], ",")
	if len(hotfixes) > 1 {
		for _, curline = range hotfixes[1:] {
			hotfix = strings.TrimSpace(strings.Split(curline, ": ")[1])
			info.OperatingSystem.Hotfixes = append(info.OperatingSystem.Hotfixes, hotfix)
		}
	}

	info.Hostname = outsplit[0]
	info.OperatingSystem.Name = outsplit[1]
	info.OperatingSystem.Version = outsplit[2]
	info.OperatingSystem.Manufacturer = outsplit[3]
	info.OperatingSystem.Domain = outsplit[28]

	return info, nil
}
