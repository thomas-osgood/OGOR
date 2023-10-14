package general

// structure designed to hold the information
// related to all the CPUs on a machine.
type AllCpuInfo struct {
	Cpus           []CpuInfo
	ProcessorCount int
	Raw            []byte
}

// structure designed to hold the basic system
// information for a given machine.
type BasicSysInfo struct {
	Hostname        string
	OperatingSystem OSInfo
	// this systemid is the productid returned
	// from the systeminfo command on windows.
	// this will be an empty string on linux.
	SystemId string
}

// structure designed to hold the information
// related to a single CPU (processor) on a machine.
type CpuInfo struct {
	ProcessorModel  string
	ProcessorSpeed  float64
	ProcessorVendor string
}

// structure designed to hold the information
// related to a system mount.
type MountInfo struct {
	Name       string
	Location   string
	Type       string
	Privileges string
}

// structure designed to hold the basic information
// about a machine's operating system.
type OSInfo struct {
	Name         string
	Version      string
	Manufacturer string
	Domain       string
	// on linux machines, this slice will be nil.
	// this is only applicable to windows machines.
	Hotfixes []string
}
