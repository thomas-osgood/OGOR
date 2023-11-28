package generics

import (
	"fmt"
	"strings"
)

// function designed to append a new int
// to an existing slice of int.
func (si SearchableIntSlice) Append(target int) (newslice SearchableIntSlice) {
	newslice = append(si, target)
	return newslice
}

// function designed to combine two SearchableIntSlices.
// this will append the entire target slice to the one it
// it being combined with.
func (si SearchableIntSlice) Combine(target SearchableIntSlice) (newslice SearchableIntSlice) {
	switch {
	case target.Length() < 1:
		newslice = si
	case si.Length() < 1:
		newslice = target
	default:
		var current int

		newslice = si
		for _, current = range si {
			newslice = newslice.Append(current)
		}
	}
	return newslice
}

// function designed to determine if the given
// int is in the int slice.
func (si SearchableIntSlice) In(target int) (err error) {
	var currentInt int
	var found bool

	if len(si) < 1 {
		return fmt.Errorf("slice is empty")
	}

	for _, currentInt = range si {
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

// function designed to get the index of the given int
// in the int slice.
func (si SearchableIntSlice) IndexOf(target int) (idx int, err error) {
	// make sure the target is in the slice that is being
	// searched. if it is not, return an error and -1.
	if err = si.In(target); err != nil {
		return -1, err
	}

	var currentIdx int
	var currentInt int

	// loop through the slice until the first instance of
	// the target int is discovered, then break the
	// loop and return the index.
	for currentIdx, currentInt = range si {
		if currentInt == target {
			idx = currentIdx
			break
		}
	}

	return idx, nil
}

// function designed to return the length of the array/slice.
// this will loop throught the slice and iterate a counter,
// returning the number of elements present in the slice.
func (si SearchableIntSlice) Length() (count int) {
	for range si {
		count++
	}
	return count
}

// function designed to return the string representation of
// the given int64 slice.
func (si SearchableIntSlice) String() string {
	var current int
	var stringSlice []string = make([]string, 0)

	for _, current = range si {
		stringSlice = append(stringSlice, fmt.Sprintf("%d", current))
	}

	return fmt.Sprintf("[%s]", strings.Join(stringSlice, ","))
}
