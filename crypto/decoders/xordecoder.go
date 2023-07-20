package decoders

import (
	"fmt"
	"strconv"
)

// function designed to loop through the ciphertext and
// xor every evenly positioned hex value with the previously
// positioned hex value.
func (xod *XORDecoder) XORPreviousPosition() (err error) {
	var curbaseten int64
	var curhex []byte
	var i int
	var pos int = 0

	for i = 0; i < len(xod.ciphertext); i += 2 {
		curhex = xod.ciphertext[i : i+2]

		// process data here
		curbaseten, err = strconv.ParseInt(string(curhex), 16, 0)
		if err != nil {
			return err
		}

		if (pos % 2) != 0 {
			curbaseten ^= int64(xod.Plaintext[pos-1])
		}

		xod.Plaintext = fmt.Sprintf("%s%c", xod.Plaintext, byte(curbaseten))

		pos += 1
	}

	return nil
}
