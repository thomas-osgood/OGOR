package decoders

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

// function designed to create and initialize a new XORDecoder object.
func NewXORDecoder(options ...XORDecoderOptsFunc) (decoder *XORDecoder, err error) {
	var optionstruct XORDecoderOptions = XORDecoderOptions{}

	decoder = &XORDecoder{}

	// loop through options and configure new decoder.
	for _, fn := range options {
		err = fn(&optionstruct)
		if err != nil {
			return nil, err
		}
	}

	// switch statement to check the ciphertext.
	switch {
	case (len(optionstruct.Ciphertext) < 1) && (len(optionstruct.Filename) < 1):
		return nil, errors.New("must specify file or ciphertext for XORDecoder")
	case (len(optionstruct.Ciphertext) > 0) && (len(optionstruct.Filename) > 0):
		return nil, errors.New("cannot specify both filename and ciphertext")
	case (len(optionstruct.Filename) > 0):
		decoder.ciphertext, err = readCiphertextFile(optionstruct.Filename)
		if err != nil {
			return nil, err
		}
	default:
		decoder.ciphertext = optionstruct.Ciphertext
	}

	// switch statement to check the encryption key.
	switch {
	case (optionstruct.EncryptionKey == nil) && (len(optionstruct.EncryptionKeyFile) < 1):
		decoder.EncryptionKey = decoder.ciphertext
	case len(optionstruct.EncryptionKeyFile) > 0:
		decoder.EncryptionKey, err = readCiphertextFile(optionstruct.EncryptionKeyFile)
		if err != nil {
			return nil, err
		}
	case optionstruct.EncryptionKey != nil:
		decoder.EncryptionKey = optionstruct.EncryptionKey
	default:
		decoder.EncryptionKey = optionstruct.EncryptionKey
	}

	return decoder, nil
}

// function designed to read a file, extract the ciphertext from it and return it.
func readCiphertextFile(filename string) (ciphertext []byte, err error) {
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

// xordecoderoptsfunc designed to specify the ciphertext to
// use when attempting to decode.
func WithCiphertext(ciphertext []byte) XORDecoderOptsFunc {
	return func(xo *XORDecoderOptions) error {
		xo.Ciphertext = ciphertext
		return nil
	}
}

// xordecoderoptsfunc designed to specify the name of the file
// in which the ciphertext is stored.
func WithFile(filename string) XORDecoderOptsFunc {
	return func(xo *XORDecoderOptions) error {
		xo.Filename = filename
		return nil
	}
}

// xordecoderoptsfunc designed to specify the encryption key
// to use.
func WithKey(keybytes []byte) XORDecoderOptsFunc {
	return func(xo *XORDecoderOptions) error {
		xo.EncryptionKey = keybytes
		return nil
	}
}

// xordecoderoptsfunc designed to specify the encryption key
// file to read the encryption key from.
func WithKeyFile(filename string) XORDecoderOptsFunc {
	return func(xo *XORDecoderOptions) error {
		xo.EncryptionKeyFile = filename
		return nil
	}
}
