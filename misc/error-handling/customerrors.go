package errorhandling

import "errors"

// function defining the string message that will
// be attached to a NotFound error.
func (nf *NotFoundError) Error() string {
	return "object not found"
}

// function designed to determine if a given error
// is of the NotFoundError type.
func (nf *NotFoundError) Is(err error) bool {
	var as *NotFoundError = new(NotFoundError)

	if err == nil {
		panic("Is comparison: err cannot be nil")
	}

	if !errors.As(err, &as) {
		return false
	}

	return true
}
