package output

import "sync"

// function designed to create and return a new
// formatter object that can be used later on.
func NewFormatter() (obj *Formatter, err error) {
	obj = &Formatter{}
	return obj, nil
}

// function designed to create and return a new
// outputter object that can be used later on.
func NewOutputter() (obj *Outputter, err error) {
	var mut *sync.Mutex = new(sync.Mutex)
	obj = &Outputter{Mutex: mut}
	return obj, nil
}
