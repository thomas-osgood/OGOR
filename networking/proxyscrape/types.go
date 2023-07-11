package proxyscrape

// function alias used to set the values of
// a ScraperOptions struct. this is used in
// the NewProxyScraper function.
type OptsFunc func(*ScraperOptions) error
