package generators

// structure designed to outline a wordlist generator.
// this will be able to read a specified wordlist and
// return its contents line-by-line. this is useful in
// enumeration and brute-forcing.
type WordlistGenerator struct {

	// file to use as the wordlist. this will be read
	// line-by-line in the ReadWordlist function.
	Wordlist string

	// channel to pass the contents of the worlist on.
	CommsChan chan string

	// flag to indicate to stop reading the wordlist.
	// this will cause the ReadWordlist function to
	// break the Scan loop and return, closing the
	// comms channel.
	StopRead bool
}
