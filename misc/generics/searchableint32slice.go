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

// function designed to combine two SearchableInt32Slices.
// this will append the entire target slice to the one it
// it being combined with.
func (si32 SearchableInt32Slice) Combine(target SearchableInt32Slice) (newslice SearchableInt32Slice) {
	switch {
	case target.Length() < 1:
		newslice = si32
	case si32.Length() < 1:
		newslice = target
	default:
		var current int32

		newslice = si32
		for _, current = range target {
			newslice = newslice.Append(current)
		}
	}
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

// function designed to return the length of the array/slice.
// this will loop throught the slice and iterate a counter,
// returning the number of elements present in the slice.
func (si32 SearchableInt32Slice) Length() (count int) {
	for range si32 {
		count++
	}
	return count
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
