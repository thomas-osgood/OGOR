package generics

import (
	"fmt"
)

// function designed to append a string to the string.
func (ss SearchableString) Append(target string) (newslice SearchableString) {
	return SearchableString(fmt.Sprintf("%s%s", ss.String(), target))
}

// function designed to append a char to the string.
func (ss SearchableString) AppendChar(target rune) (newslice SearchableString) {
	return SearchableString(fmt.Sprintf("%s%c", ss.String(), target))
}

// function designed to search through a searchablestring
// and determine if the target rune (char) is in it.
func (ss SearchableString) In(target string) (err error) {
	var current rune
	var found bool

	// make sure only a single character is passed in.
	if len(target) != 1 {
		return fmt.Errorf("can only search for a single char.")
	}

	// loop over each character in the string and compare
	// it to the target char. if a match is found, set the
	// found flag and break the loop.
	for _, current = range ss {
		if fmt.Sprintf("%c", current) == target {
			found = true
			break
		}
	}

	// if the char was not found within the string, return
	// an error stating as much.
	if !found {
		return fmt.Errorf("%s not found in %s", target, ss)
	}

	return nil
}

// function designed to find the index of a given char.
func (ss SearchableString) IndexOf(target string) (idx int, err error) {
	// make sure the target is in the string before
	// looping through to grab the index.
	if err = ss.In(target); err != nil {
		return -1, err
	}

	var currentIdx int
	var currentChr rune

	for currentIdx, currentChr = range ss {
		if fmt.Sprintf("%c", currentChr) == target {
			idx = currentIdx
			break
		}
	}

	return idx, nil
}

func (ss SearchableString) String() string {
	return string(ss)
}