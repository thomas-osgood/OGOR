package replicator

import (
	"os"
	"syscall"
)

// function designed to replicate the currently running
// process. this will spawn a new process and "Release"
// it, allowing it to continue running even after the
// current process has exited.
//
// this attempts to mimic the FORK command found in unix
// systems on both UNIX and non-UNIX systems (ie: this
// should work on Windows, too).
func Replicate(args []string) (err error) {
	var attr os.ProcAttr
	var currentProc string
	var procargs []string = make([]string, 0)
	var spawned *os.Process
	var sysproc *syscall.SysProcAttr

	// get the full name of the executable that spawned
	// the currently running process. this will be used
	// to spawn the clone later on in this function.
	currentProc, err = os.Executable()
	if err != nil {
		return err
	}

	// setup the arguments to the new process. argv[0]
	// is always the process name.
	procargs = append(procargs, currentProc)
	procargs = append(procargs, args...)

	sysproc = &syscall.SysProcAttr{}

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
	spawned, err = os.StartProcess(currentProc, procargs, &attr)
	if err != nil {
		return err
	}

	// Release (aka "detach") the process so it
	// can keep running after this one exits.
	err = spawned.Release()
	if err != nil {
		return err
	}

	return nil
}
