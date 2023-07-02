package generators

import (
	"bufio"
	"os"
)

// function designed to go through the WordlistGenerator's
// designated wordlist line-by-line and feed each line to
// the comms channel.
//
// when this function returns, the comms channel will close.
func (w *WordlistGenerator) ReadWordlist() (err error) {
	defer close(w.CommsChan)

	var fptr *os.File
	var line string
	var scanner *bufio.Scanner

	fptr, err = os.Open(w.Wordlist)
	if err != nil {
		return err
	}
	defer fptr.Close()

	scanner = bufio.NewScanner(fptr)

	// read file line-by-line.
	for scanner.Scan() {
		if w.StopRead {
			break
		}

		line = scanner.Text()
		w.CommsChan <- line
	}

	return nil
}
