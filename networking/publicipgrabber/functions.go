package publicipgrabber

import "net/http"

// function designed to create and initialize a new
// PubliIPGrabber object. the user can pass in option
// functions to change the configuration.
func NewPublicIPGrabber(optfuncs ...PublicIPGrabberOptFunc) (grabber *PublicIPGrabber, err error) {
	var fn PublicIPGrabberOptFunc
	var options PublicIPGrabberOptions = PublicIPGrabberOptions{Client: http.DefaultClient}

	grabber = &PublicIPGrabber{}

	// loop through and set configuration options.
	for _, fn = range optfuncs {
		err = fn(&options)
		if err != nil {
			return nil, err
		}
	}

	grabber.client = options.Client

	return grabber, nil
}

// function designed to set the PublicIPGrabberOptions client.
func WithClient(client *http.Client) PublicIPGrabberOptFunc {
	return func(pio *PublicIPGrabberOptions) error {
		pio.Client = client
		return nil
	}
}
