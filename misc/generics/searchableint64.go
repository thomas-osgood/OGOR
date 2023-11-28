package generics

import (
	"fmt"
	"strconv"
)

// function designed to append a new int64
// to an existing slice of int64.
func (si64 SearchableInt64) Append(target int64) (newslice SearchableInt64) {
	var err error
	var newNumber int
	var newNumberString string = fmt.Sprintf("%s%c", si64.String(), target)

	newNumber, err = strconv.Atoi(newNumberString)
	if err != nil {
		newNumber = int(si64)
	}

	newslice = SearchableInt64(int64(newNumber))

	return newslice
}

// function designed to determine if the given
// int64 is in the int64.
func (si64 SearchableInt64) In(target int64) (err error) {
	var currentInt rune
	var found bool
	var searchString string = si64.String()

	if len(searchString) < 1 {
		return fmt.Errorf("number is empty")
	}

	for _, currentInt = range searchString {
		if currentInt == int32(target) {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%d not found in the number", target)
	}

	return nil
}

// function designed to get the index of the given int64
// in the int64.
func (si64 SearchableInt64) IndexOf(target int64) (idx int, err error) {
	// make sure the target is in the slice that is being
	// searched. if it is not, return an error and -1.
	if err = si64.In(target); err != nil {
		return -1, err
	}

	var currentIdx int
	var currentInt rune
	var searchString string = si64.String()

	// loop through the slice until the first instance of
	// the target int64 is discovered, then break the
	// loop and return the index.
	for currentIdx, currentInt = range searchString {
		if currentInt == int32(target) {
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
func (si64 SearchableInt64) Length() (count int) {
	for range si64.String() {
		count++
	}
	return count
}

// function designed to return the string representation of
// the given int64 slice.
func (si64 SearchableInt64) String() string {
	return fmt.Sprintf("%d", si64)
}
