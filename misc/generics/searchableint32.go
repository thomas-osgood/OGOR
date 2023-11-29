package generics

import (
	"fmt"
	"strconv"
)

// function designed to append a new int32
// to an existing slice of int32.
func (si32 SearchableInt32) Append(target int32) (newslice SearchableInt32) {
	var err error
	var newNumber int
	var newNumberString string = fmt.Sprintf("%s%c", si32.String(), target)

	newNumber, err = strconv.Atoi(newNumberString)
	if err != nil {
		newNumber = int(si32)
	}

	newslice = SearchableInt32(int32(newNumber))

	return newslice
}

// function designed to reset the object to
// the state it is when initialized. this will
// set numbers to zero.
func (si32 SearchableInt32) Clear() SearchableInt32 {
	return 0
}

// function designed to determine if the given
// int32 is in the int32.
func (si32 SearchableInt32) In(target int32) (err error) {
	var currentInt rune
	var found bool
	var searchString string = si32.String()

	if len(searchString) < 1 {
		return fmt.Errorf("number is empty")
	}

	for _, currentInt = range searchString {
		if currentInt == target {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%d not found in the number", target)
	}

	return nil
}

// function designed to get the index of the given int32
// in the int32.
func (si32 SearchableInt32) IndexOf(target int32) (idx int, err error) {
	// make sure the target is in the slice that is being
	// searched. if it is not, return an error and -1.
	if err = si32.In(target); err != nil {
		return -1, err
	}

	var currentIdx int
	var currentInt rune
	var searchString string = si32.String()

	// loop through the slice until the first instance of
	// the target int32 is discovered, then break the
	// loop and return the index.
	for currentIdx, currentInt = range searchString {
		if currentInt == target {
			idx = currentIdx
			break
		}
	}

	return idx, nil
}

// function designed to return the length of the object.
// this will loop throught the object and iterate a counter,
// returning the number of elements present in the string
// representation of the object.
func (si32 SearchableInt32) Length() (count int) {
	for range si32.String() {
		count++
	}
	return count
}

// function designed to return the string representation of
// the given searchable int32.
func (si32 SearchableInt32) String() string {
	return fmt.Sprintf("%d", si32)
}
