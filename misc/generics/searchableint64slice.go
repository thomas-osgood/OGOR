package generics

import (
	"fmt"
	"strings"
)

// function designed to append a new int64
// to an existing slice of int64.
func (si64 SearchableInt64Slice) Append(target int64) (newslice SearchableInt64Slice) {
	newslice = append(si64, target)
	return newslice
}

// function designed to determine if the given
// int64 is in the int64 slice.
func (si64 SearchableInt64Slice) In(target int64) (err error) {
	var currentInt int64
	var found bool

	if len(si64) < 1 {
		return fmt.Errorf("slice is empty")
	}

	for _, currentInt = range si64 {
		if currentInt == target {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%d not found in slice", target)
	}

	return nil
}

// function designed to get the index of the given int64
// in the int64 slice.
func (si64 SearchableInt64Slice) IndexOf(target int64) (idx int, err error) {
	// make sure the target is in the slice that is being
	// searched. if it is not, return an error and -1.
	if err = si64.In(target); err != nil {
		return -1, err
	}

	var currentIdx int
	var currentInt int64

	// loop through the slice until the first instance of
	// the target int64 is discovered, then break the
	// loop and return the index.
	for currentIdx, currentInt = range si64 {
		if currentInt == target {
			idx = currentIdx
			break
		}
	}

	return idx, nil
}

// function designed to return the string representation of
// the given int64 slice.
func (si64 SearchableInt64Slice) String() string {
	var current int64
	var stringSlice []string = make([]string, 0)

	for _, current = range si64 {
		stringSlice = append(stringSlice, fmt.Sprintf("%d", current))
	}

	return fmt.Sprintf("[%s]", strings.Join(stringSlice, ","))
}
