package proxyscrape

// set anonymity level to "all".
const ANONYMITY_ALL int = 0

// set anonymity level to "anonymous".
const ANONYMITY_ANONYMOUS int = 2

// set anonymity level to "elite".
const ANONYMITY_ELITE int = 3

// set anonymity level to "transparent".
const ANONYMITY_TRANSPARENT int = 1

// this is the API url for proxyscrape v2.
const BASE_URL string = "https://api.proxyscrape.com/v2"

// set the ssl option to "all"
const HTTPS_ALL int = 0

// set the ssl option to "yes"
const HTTPS_YES int = 1

// set the ssl option to "no"
const HTTPS_NO int = 2

// maximum allowed timeout value that ProxyScraper can get.
const MAX_TIMEOUT int = 10000

// set protocol to all
const PROTO_ALL int = 0

// set protocol to http
const PROTO_HTTP int = 1

// set protocol to socks4
const PROTO_SOCKS4 int = 2

// set protocol to socks5
const PROTO_SOCKS5 int = 3
