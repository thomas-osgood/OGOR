package forwardproxy

import "time"

// constant defining the default port the forwarder
// will be served on.
const DEFAULT_PORTNO int = 10000

// constant defining the default timeout for a request.
const DEFAULT_TIMEOUT time.Duration = 10 * time.Second
