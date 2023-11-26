package generics

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

// function designed to append a string to the string slice.
func (ss SearchableStringSlice) Append(target string) (newslice SearchableStringSlice) {
	return append(ss, target)
}

// function designed to append a char to the string slice.
func (ss SearchableStringSlice) AppendChar(target rune) (newslice SearchableStringSlice) {
	return append(ss, fmt.Sprintf("%c", target))
}

// function designed to search through a searchablestring
// and determine if the target string is in it.
func (ss SearchableStringSlice) In(target string) (err error) {
	var current string
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
		return fmt.Errorf("%s not found in %s", target, ss)
	}

	return nil
}

// function designed to find the index of a given string.
func (ss SearchableStringSlice) IndexOf(target string) (idx int, err error) {
	// make sure the target is in the string before
	// looping through to grab the index.
	if err = ss.In(target); err != nil {
		return -1, err
	}

	var currentIdx int
	var currentChr string

	for currentIdx, currentChr = range ss {
		if currentChr == target {
			idx = currentIdx
			break
		}
	}

	return idx, nil
}

// function designed to sort the SearchableStringSlice
// alphabetically.
func (ss SearchableStringSlice) Sort() SearchableStringSlice {
	slices.Sort(ss)
	return ss
}

// function designed to return the string representation
// of the SearchableStringSlice.
func (ss SearchableStringSlice) String() string {
	return fmt.Sprintf("[%s]", strings.Join(ss, ","))
}
