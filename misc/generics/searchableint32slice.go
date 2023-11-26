package generics

import (
	"fmt"
	"strings"
)

// function designed to append a new int32
// to an existing slice of int32.
func (si32 SearchableInt32Slice) Append(target int32) (newslice SearchableInt32Slice) {
	newslice = append(si32, target)
	return newslice
}

// function designed to determine if the given
// int32 is in the int32 slice.
func (si32 SearchableInt32Slice) In(target int32) (err error) {
	var currentInt int32
	var found bool

	if len(si32) < 1 {
		return fmt.Errorf("slice is empty")
	}

	for _, currentInt = range si32 {
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

// function designed to get the index of the given int32
// in the int32 slice.
func (si32 SearchableInt32Slice) IndexOf(target int32) (idx int, err error) {
	// make sure the target is in the slice that is being
	// searched. if it is not, return an error and -1.
	if err = si32.In(target); err != nil {
		return -1, err
	}

	var currentIdx int
	var currentInt int32

	// loop through the slice until the first instance of
	// the target int32 is discovered, then break the
	// loop and return the index.
	for currentIdx, currentInt = range si32 {
		if currentInt == target {
			idx = currentIdx
			break
		}
	}

	return idx, nil
}

// function designed to return the string representation of
// the given int32 slice.
func (si32 SearchableInt32Slice) String() string {
	var current int32
	var stringSlice []string = make([]string, 0)

	for _, current = range si32 {
		stringSlice = append(stringSlice, fmt.Sprintf("%d", current))
	}

	return fmt.Sprintf("[%s]", strings.Join(stringSlice, ","))
}
