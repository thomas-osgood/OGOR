package proxyscrape

// this represents the data that will be returned by
// proxyscrape when the API is hit.
type ProxyList struct {
	Proxies []string
}

// this is the object that will be used to contact the
// API and pull down and store the proxies list.
type ProxyScraper struct {

	// this represents the anonymity level desired
	// for the proxy. this only applies to HTTP(S) proxies.
	// the values for this are defined as constants.
	anonymity int

	// this represents the country that the proxy is
	// located in. it is the 2 letter country code.
	// to ignore country, this should be set to all.
	country string

	// the proxy protocol to filter on. if set to all
	// this will grab all protocols.
	protocol int

	// list of proxies returned by proxyscrape.
	Proxies ProxyList

	// specifies whether the proxy can use HTTPS. the
	// only options are "yes", "no", and "all". these
	// are defined as constants.
	ssl int

	// specified the max timeout for the proxy (limit 10K).
	timeout int
}

type ScraperOptions struct {

	// this represents the anonymity level desired
	// for the proxy. this only applies to HTTP(S) proxies.
	// the values for this are defined as constants.
	Anonymity int

	// this represents the country that the proxy is
	// located in. it is the 2 letter country code.
	// to ignore country, this should be set to all.
	Country string

	// the proxy protocol to filter on. if set to all
	// this will grab all protocols.
	Protocol int

	// specifies whether the proxy can use HTTPS. the
	// only options are "yes", "no", and "all". these
	// are defined as constants.
	SSL int

	// specified the max timeout for the proxy (limit 10K).
	Timeout int
}
