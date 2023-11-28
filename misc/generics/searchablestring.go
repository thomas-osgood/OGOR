package generics

import (
	"fmt"
)

// function designed to append a string to the string.
func (ss SearchableString) Append(target rune) (newslice SearchableString) {
	return SearchableString(fmt.Sprintf("%s%c", ss.String(), target))
}

// function designed to append a string to the string.
func (ss SearchableString) AppendString(target string) (newslice SearchableString) {
	return SearchableString(fmt.Sprintf("%s%s", ss.String(), target))
}

// function designed to append a char to the string.
func (ss SearchableString) AppendChar(target rune) (newslice SearchableString) {
	return SearchableString(fmt.Sprintf("%s%c", ss.String(), target))
}

// function designed to search through a searchablestring
// and determine if the target rune (char) is in it.
func (ss SearchableString) In(target rune) (err error) {
	var current rune
	var found bool

	// loop over each character in the string and compare
	// it to the target char. if a match is found, set the
	// found flag and break the loop.
	for _, current = range ss {
		if current == target {
			found = true
			break
		}
	}

	// if the char was not found within the string, return
	// an error stating as much.
	if !found {
		return fmt.Errorf("%c not found in %s", target, ss)
	}

	return nil
}

// function designed to find the index of a given char.
func (ss SearchableString) IndexOf(target rune) (idx int, err error) {
	// make sure the target is in the string before
	// looping through to grab the index.
	if err = ss.In(target); err != nil {
		return -1, err
	}

	var currentIdx int
	var currentChr rune

	for currentIdx, currentChr = range ss {
		if currentChr == target {
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
func (ss SearchableString) Length() (count int) {
	for range ss {
		count++
	}
	return count
}

func (ss SearchableString) String() string {
	return string(ss)
}
