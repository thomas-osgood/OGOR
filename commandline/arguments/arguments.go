// this package defines and initializes custom types
// that can be used with the Flag package. these types
// can be called using the flag.Var() function.
//
// Common Types:
//
// - string array: comma-separated list of strings that can be passed in.
//
// - int array: comma-separated list of ints that can be passed in.
package arguments

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// function designed to set the values of an integer array object. this
// will take in a string, process it, convert its values into integers,
// and append them to the new integer array.
func (i *IntArray) Set(value string) (err error) {
	var current string
	var curint int
	var splitval []string

	// split the argument on commas
	splitval = strings.Split(value, ",")

	if len(splitval) < 1 {
		return errors.New("no value passed to argument")
	}

	// build the integer slice. if there is an issue
	// converting the current string to an integer,
	// ignore it and move onto the next element.
	for _, current = range splitval {
		curint, err = strconv.Atoi(current)
		if err != nil {
			continue
		}
		*i = append(*i, curint)
	}

	return nil
}

// function designed to return the string representatino of the custom
// IntArray object. this will process the values in the integer array and
// pass back a single string.
func (i *IntArray) String() (val string) {
	var current int

	// build the string representation and trim all
	// leading and trailing spaces.
	for _, current = range *i {
		val = fmt.Sprintf("%s %d", val, current)
	}
	val = strings.Trim(val, " ")

	return val
}

// function designed to set the values of the string array object. this
// will take in a string, process it, and convert it into a string array.
func (s *StringArray) Set(value string) (err error) {
	var current string
	var splitval []string

	// split the argument on commas
	splitval = strings.Split(value, ",")

	if len(splitval) < 1 {
		return errors.New("no value passed to argument")
	}

	// build the string slice
	for _, current = range splitval {
		*s = append(*s, current)
	}

	return nil
}

// function designed to return the string representation of the custom
// StringArray object. this will process the values in the string array
// and pass back a single string.
func (s *StringArray) String() (val string) {
	var current string

	for _, current = range *s {
		val = fmt.Sprintf("%s %s", val, current)
	}
	val = strings.Trim(val, " ")
	return val
}

