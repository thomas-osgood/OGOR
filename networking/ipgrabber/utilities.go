package ipgrabber

// function designed to create a new IPGrabber, populate both
// slices and return the pointer to the user. if an error occurs
// nil and the error will be returned, otherwise the pointer and
// nil will be returned.
func NewGrabber() (grabber *IPGrabber, err error) {
	grabber = &IPGrabber{}

	// populate interfaces slice
	err = grabber.GrabInterfaces()
	if err != nil {
		return nil, err
	}

	// populate IPs slice
	err = grabber.GrabIPs()
	if err != nil {
		return nil, err
	}

	return grabber, nil
}
