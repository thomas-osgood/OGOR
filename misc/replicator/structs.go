package replicator

type Replicator struct {
	// name of the program (full path) to replicate.
	progname string
	// comamnd-line arguments to pass to the process.
	args []string
}

type ReplicatorOpts struct {
	// name of the program (full path) to replicate.
	ProgramName string
	// command-line arguments to pass to the process.
	Args []string
}
