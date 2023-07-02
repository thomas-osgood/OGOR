package generators

import "os"

// function designed to create and return a new instance of a
// wordlist generator. this will create a new comms channel and
// use the wordlist specified in the arguments.
func NewWordlistGenerator(wordlist string) (generator *WordlistGenerator, err error) {
	var fptr *os.File

	// confirm wordlist file exists before moving on.
	fptr, err = os.Open(wordlist)
	if err != nil {
		return nil, err
	}
	fptr.Close()

	generator = &WordlistGenerator{}

	generator.Wordlist = wordlist
	generator.CommsChan = make(chan string)
	generator.StopRead = false

	return generator, nil
}
