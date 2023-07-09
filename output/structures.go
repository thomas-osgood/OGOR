package output

import "sync"

// object designed to output formatted strings uising ANSI colors.
type Outputter struct {
	// built-in mutext used to avoid collisions when outputting
	// data using threading.
	Mutex *sync.Mutex
}

// object designed to format output using ANSI colors.
// similar to Outputter.
type Formatter struct{}
