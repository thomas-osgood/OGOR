// package designed to hold common functions that are related to
// handling errors encountered during execution.
package errorhandling

import (
	"fmt"
	"os"
)

// function designed to compress the normal error check down to one
// line and make handling critical errors take less space. this will
// print out the error and exit the program with a return code of 1
// if the error passed into it is not nil.
func HandleCritical(err error) {
	if err != nil {
		fmt.Printf("[-] %s\n", err.Error())
		os.Exit(1)
	}
	return
}
