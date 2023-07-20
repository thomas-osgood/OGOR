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
