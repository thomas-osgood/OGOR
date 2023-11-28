package generics

import (
	"fmt"
	"strconv"
)

// function designed to append a new int
// to an existing slice of int.
func (si SearchableInt) Append(target int) (newslice SearchableInt) {
	var err error
	var newNumber int
	var newNumberString string = fmt.Sprintf("%s%c", si.String(), target)

	newNumber, err = strconv.Atoi(newNumberString)
	if err != nil {
		newNumber = int(si)
	}

	newslice = SearchableInt(int(newNumber))

	return newslice
}

// function designed to determine if the given
// int is in the int.
func (si SearchableInt) In(target int) (err error) {
	var currentInt rune
	var found bool
	var searchString string = si.String()

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

// function designed to get the index of the given int
// in the int.
func (si SearchableInt) IndexOf(target int) (idx int, err error) {
	// make sure the target is in the slice that is being
	// searched. if it is not, return an error and -1.
	if err = si.In(target); err != nil {
		return -1, err
	}

	var currentIdx int
	var currentInt rune
	var searchString string = si.String()

	// loop through the slice until the first instance of
	// the target int is discovered, then break the
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
func (si SearchableInt) Length() (count int) {
	for range si.String() {
		count++
	}
	return count
}

// function designed to return the string representation of
// the given searchable int.
func (si SearchableInt) String() string {
	return fmt.Sprintf("%d", si)
}
