package publicipgrabber

// type alias defining the function structure
// that will be used to set the configuration
// options for a PublicIPGrabberOptions object.
type PublicIPGrabberOptFunc func(*PublicIPGrabberOptions) error
