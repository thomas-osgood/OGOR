package encoders

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

// function designed to loop through the ciphertext and
// xor every evenly positioned hex value with the previously
// positioned hex value.
func (xoe *XOREncoder) XORPreviousPosition() (err error) {
	var curbaseten int64
	var curhex []byte
	var i int
	var plainhex string = hex.EncodeToString([]byte(xoe.plaintext))
	var pos int = 0

	xoe.Ciphertext = ""

	for i = 0; i < len(plainhex); i += 2 {
		curhex = []byte(plainhex[i : i+2])

		// process data here
		curbaseten, err = strconv.ParseInt(string(curhex), 16, 0)
		if err != nil {
			return err
		}

		if (pos % 2) != 0 {
			curbaseten ^= int64(xoe.plaintext[pos-1])
		}

		xoe.Ciphertext = fmt.Sprintf("%s%02x", string(xoe.Ciphertext), byte(curbaseten))

		pos += 1
	}

	return nil
}

// function designed to loop through the plaintext and encrypt each
// character using the key provided.
func (xoe *XOREncoder) XOREncryptWithKey() (err error) {
	var cipherval int
	var i int = 0
	var keylen int = len(string(xoe.key))
	var keypos int
	var keyval int
	var offset int = xoe.keyoffset % keylen
	var plainval int

	keypos = offset

	xoe.Ciphertext = ""

	for i = 0; i < len(xoe.plaintext); i++ {
		keyval = int(byte(xoe.key[keypos]))
		plainval = int(byte(xoe.plaintext[i]))

		cipherval = plainval ^ keyval

		xoe.Ciphertext = fmt.Sprintf("%s%02x", xoe.Ciphertext, byte(cipherval))

		keypos = (keypos + 1) % keylen

	}

	return nil
}
