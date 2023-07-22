package encoders

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

// function designed to create and initialize a new XORDecoder object.
func NewXOREncoder(options ...XOREncoderOptsFunc) (encoder *XOREncoder, err error) {
	var optionstruct XOREncoderOptions = XOREncoderOptions{}

	encoder = &XOREncoder{}

	// loop through options and configure new encoder.
	for _, fn := range options {
		err = fn(&optionstruct)
		if err != nil {
			return nil, err
		}
	}

	// switch statement to check the plaintext.
	switch {
	case (len(optionstruct.Plaintext) < 1) && (len(optionstruct.PlaintextFile) < 1):
		return nil, errors.New("must specify file or plaintext for XOREncoder")
	case (len(optionstruct.Plaintext) > 0) && (len(optionstruct.PlaintextFile) > 0):
		return nil, errors.New("cannot specify both filename and plaintext")
	case (len(optionstruct.PlaintextFile) > 0):
		encoder.plaintext, err = readPlaintextFile(optionstruct.PlaintextFile)
		if err != nil {
			return nil, err
		}
	default:
		encoder.plaintext = optionstruct.Plaintext
	}

	// switch statement to check the encryption key.
	switch {
	case (optionstruct.Key == nil) && (len(optionstruct.KeyFile) < 1):
		encoder.key = nil
	case (optionstruct.Key != nil) && (len(optionstruct.KeyFile) > 0):
		return nil, errors.New("cannot specify both Key and KeyFile")
	case len(optionstruct.KeyFile) > 0:
		encoder.key, err = readKeyFile(optionstruct.KeyFile)
		if err != nil {
			return nil, err
		}
	default:
		encoder.key = optionstruct.Key
	}

	// set encryption key offset
	encoder.keyoffset = optionstruct.KeyOffset

	return encoder, nil
}

// function designed to read a file, extract the ciphertext from it and return it.
func readKeyFile(filename string) (ciphertext []byte, err error) {
	var fptr *os.File

	fptr, err = os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fptr.Close()

	ciphertext, err = ioutil.ReadAll(fptr)
	if err != nil {
		return nil, err
	}

	ciphertext = []byte(strings.Trim(string(ciphertext), " \n\r\t"))

	if (len(ciphertext) < 1) || (ciphertext == nil) {
		return nil, errors.New("cipher file empty")
	}

	return ciphertext, nil
}

// function designed to read a file, extract the plaintext from it and return it.
func readPlaintextFile(filename string) (plaintext string, err error) {
	var contents []byte
	var fptr *os.File

	fptr, err = os.Open(filename)
	if err != nil {
		return "", err
	}
	defer fptr.Close()

	contents, err = ioutil.ReadAll(fptr)
	if err != nil {
		return "", err
	}

	plaintext = strings.Trim(string(contents), " \n\r\t")

	if len(plaintext) < 1 {
		return "", errors.New("plaintext file empty")
	}

	return plaintext, nil
}

// xorencoderoptsfunc designed to specify the name of the file
// in which the plaintext is stored.
func WithFile(filename string) XOREncoderOptsFunc {
	return func(xo *XOREncoderOptions) error {
		xo.PlaintextFile = filename
		return nil
	}
}

// xorencoderoptsfunc designed to specify the encryption key
// to use.
func WithKey(keybytes []byte) XOREncoderOptsFunc {
	return func(xo *XOREncoderOptions) error {
		xo.Key = keybytes
		return nil
	}
}

// xorencoderoptsfunc designed to specify the encryption key
// file to read the encryption key from.
func WithKeyFile(filename string) XOREncoderOptsFunc {
	return func(xo *XOREncoderOptions) error {
		xo.KeyFile = filename
		return nil
	}
}

// xorencoderoptsfunc designed to specify the key offset
// to use when encrypting plaintext.
func WithKeyOffset(offset int) XOREncoderOptsFunc {
	return func(xo *XOREncoderOptions) error {
		xo.KeyOffset = offset
		return nil
	}
}

// xorencoderoptsfunc designed to specify the plaintext the
// user wants to encrypt.
func WithPlaintext(plaintext string) XOREncoderOptsFunc {
	return func(xo *XOREncoderOptions) error {
		xo.Plaintext = plaintext
		return nil
	}
}
