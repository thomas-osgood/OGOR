//go:build windows
// +build windows

package replicator

// function designed to replicate the currently running
// process. this will spawn a new, detached, process,
// allowing it to continue running even after the
// current process has exited.
func (r *Replicator) Replicate() (err error) {
	var argv *uint16
	var procinfo syscall.ProcessInformation
	var startupinfo syscall.StartupInfo

	// get the full name of the executable that spawned
	// the currently running process. this will be used
	// to spawn the clone later on in this function.
	currentProc, err = os.Executable()
	if err != nil {
		return err
	}
	currentProc = fmt.Sprintf("%s %s", r.progname, strings.Join(r.args, " "))

	argv = syscall.StringToUTF16Ptr(currentProc)

	// make Win32API call to create the new, detached, process. this
	// will spawn the process with the child flag set, opening a new
	// terminal and allowing the parent to exit without interfering
	// with the child's logic.
	return syscall.CreateProcess(
		nil,              // name of the child process to spawn. this can be nil.
		argv,             // command-line arguments to pass to the process. argv[0] should be the binary's full path.
		nil,              // security attributes. this can be nil.
		nil,              // thread security attributes. this can be nil.
		false,            // bool flag to indicate whether to inherid from parent process. for detached this should be false.
		CREATE_NO_WINDOW, // creation flags. see above for more info.
		nil,              // use parent's environment block. this can be nil.
		nil,              // use parent's starting directory. this can be nil.
		&startupinfo,     // pointer to startupinfo structure.
		&procinfo,        // pointer to processinformation structure.
	)
}
