package decoders

import (
	"encoding/hex"
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

// function designed to decrypt ciphertext using a key designated
// by the user.
func (xod *XORDecoder) XORDecryptWithKey(offset int) (err error) {
	var curbaseten int64
	var curletter byte
	var i int = 0
	var keylen int = len(string(xod.EncryptionKey))
	var keypos int = offset
	var keystring []byte

	xod.Plaintext = ""

	keystring, err = hex.DecodeString(string(xod.EncryptionKey))
	if err != nil {
		return err
	}
	keylen = len(keystring)

	for i = 0; i < len(xod.ciphertext); i += 2 {
		curbaseten, err = strconv.ParseInt(string(xod.ciphertext[i:i+2]), 16, 0)
		if err != nil {
			return err
		}
		curletter = byte(curbaseten ^ int64(keystring[keypos]))
		xod.Plaintext = fmt.Sprintf("%s%c", xod.Plaintext, curletter)

		keypos = (keypos + 1) % keylen
	}

	return nil
}
