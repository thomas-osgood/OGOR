//go:build !windows
// +build !windows

package replicator

import (
	"os"
	"syscall"
)

// function designed to replicate the currently running
// process. this will spawn a new, detached, process,
// allowing it to continue running even after the
// current process has exited.
func (r *Replicator) Replicate() (err error) {
	var attr os.ProcAttr
	var procargs []string = make([]string, 0)
	var spawned *os.Process
	var sysproc *syscall.SysProcAttr

	// setup the arguments to the new process. argv[0]
	// is always the process name.
	procargs = append(procargs, r.progname)
	procargs = append(procargs, r.args...)

	// set the process attributes to make sure the process
	// starts in its own process group and not the process
	// group of the process that spawned it.
	sysproc = &syscall.SysProcAttr{
		Setpgid: true,
		Noctty:  true,
		Pgid:    0,
	}

	attr = os.ProcAttr{
		Dir: ".",
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			nil,
			nil,
		},
		Sys: sysproc,
	}

	// spawn the new process with the command-line arguments
	// provided via the args parameter of this function.
	spawned, err = os.StartProcess(r.progname, procargs, &attr)
	if err != nil {
		return err
	}

	// Release (aka "detach") the process so it
	// can keep running after this one exits.
	return spawned.Release()
}
