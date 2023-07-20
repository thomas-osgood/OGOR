package decoders

// structure representing an XOR decoder. this will carry
// out all the functions related to decoding/decrypting a
// message that has been encoded using XOR methods.
type XORDecoder struct {

	// this is the hexidecimal message to be decoded.
	ciphertext []byte

	// this is the decoded plaintext message.
	Plaintext string
}

// structure representing an OptsStruct that will hold the
// configuration options for an XORDecoder object.
type XORDecoderOptions struct {
	// if this variable is set, it holds the file in which
	// the ciphertext is saved. this file will be read and
	// the content will be loaded into an XORDecoder's
	// ciphertext variable.
	//
	// note: this cannot be set alongside Ciphertext
	Filename string

	// if this variable is set, it hodls the ciphertext that
	// the XORDecoder will decode.
	//
	// note: this cannot be set alongside Filename.
	Ciphertext []byte
}
