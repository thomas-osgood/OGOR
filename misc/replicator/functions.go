package replicator

import (
	"os"
)

// function designed to initialize and return a new replicator.
func NewReplicator(optFuncs ...RepOptFunc) (replicator *Replicator, err error) {
	var currentprog string
	var currentSetting RepOptFunc
	var defaultOptions ReplicatorOpts = ReplicatorOpts{}

	if currentprog, err = os.Executable(); err != nil {
		return nil, err
	}

	defaultOptions.ProgramName = currentprog
	defaultOptions.Args = []string{}

	for _, currentSetting = range optFuncs {
		if err = currentSetting(&defaultOptions); err != nil {
			return nil, err
		}
	}

	replicator = new(Replicator)

	replicator.progname = defaultOptions.ProgramName
	replicator.args = defaultOptions.Args

	return replicator, nil
}
