package validations

import (
	"errors"
	"fmt"
)

// function designed to validate a given network port number. if
// the port if too low or too high, this will be indicated in the
// return, othwerwise nil will be returned.
func ValidateNetworkPort(port int) (err error) {
	// make sure port number falls within range 1 - 65535.
	// indicate too low or too high.
	if port < PORT_MIN {
		return errors.New(fmt.Sprintf("port number too low. port must be with range %d to %d", PORT_MIN, PORT_MAX))
	} else if port > PORT_MAX {
		return errors.New(fmt.Sprintf("port number too high. port must be with range %d to %d", PORT_MIN, PORT_MAX))
	}
	return nil
}
