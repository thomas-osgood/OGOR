package general

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
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

	outbytes = bytes.TrimSpace(outbytes)
	if len(outbytes) < 1 {
		return nil, fmt.Errorf("no mounts discovered")
	}

	outsplit = bytes.Split(outbytes, []byte("\n"))

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

// function deisgned to read the cpuinfo file and pull
// out information of interest. this will help give a
// better understanding of the machine's CPU(s).
func GetCPUInfo() (info AllCpuInfo, err error) {
	var proccntpat string = "processor\\s+:\\s?[0-9]+"
	var procmodpat string = "model\\sname\\s+:.*"
	var procspdpat string = "cpu\\sMHz\\s+:\\s?[0-9.]+"
	var fptr *os.File
	var idx int
	var match []byte
	var matches [][]byte
	var numprocs int
	var processor []byte
	var processors [][]byte
	var rawdata []byte
	var re *regexp.Regexp
	var repm *regexp.Regexp
	var reps *regexp.Regexp
	var splitmatch [][]byte
	const target string = "/proc/cpuinfo"

	// initialize Cpus slice in AllCpuInfo struct
	info.Cpus = make([]CpuInfo, 0)

	fptr, err = os.Open(target)
	if err != nil {
		return AllCpuInfo{}, err
	}
	defer fptr.Close()

	rawdata, err = io.ReadAll(fptr)
	if err != nil {
		return AllCpuInfo{}, err
	}

	rawdata = bytes.TrimSpace(rawdata)

	info = AllCpuInfo{
		Raw: rawdata,
	}

	// setup regex to pullout the different processors
	re, err = regexp.Compile(proccntpat)
	if err != nil {
		return info, err
	}

	// pull out all processors discovered in the file
	matches = re.FindAll(rawdata, -1)
	if matches == nil {
		numprocs = 0
	} else {
		numprocs = len(matches)
	}

	info.ProcessorCount = numprocs

	// setup regex to pull out the processor model
	repm, err = regexp.Compile(procmodpat)
	if err != nil {
		return info, err
	}

	// setup regex to pull out the processor speed
	reps, err = regexp.Compile(procspdpat)
	if err != nil {
		return info, err
	}

	processors = bytes.Split(rawdata, []byte("\n\n"))
	for idx, processor = range processors {
		info.Cpus = append(info.Cpus, CpuInfo{})

		matches = repm.FindAll(processor, -1)
		if matches != nil {
			splitmatch = bytes.Split(matches[0], []byte(":"))
			match = bytes.TrimSpace(splitmatch[len(splitmatch)-1])
			info.Cpus[idx].ProcessorModel = string(match)
		}

		matches = reps.FindAll(processor, -1)
		if matches != nil {
			splitmatch = bytes.Split(matches[0], []byte(":"))
			match = bytes.TrimSpace(splitmatch[len(splitmatch)-1])
			info.Cpus[idx].ProcessorSpeed, err = strconv.ParseFloat(string(match), 64)
		}
	}

	return info, nil
}
