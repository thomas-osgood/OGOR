package general

// structure designed to hold the information
// related to all the CPUs on a machine.
type AllCpuInfo struct {
	Cpus           []CpuInfo
	ProcessorCount int
	Raw            []byte
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
