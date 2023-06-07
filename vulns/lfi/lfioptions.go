package lfichecker

// function designed to set the DoubleEncoding flag for
// a given LFIOptions object.
func UsingDoubleEncoding(opt *LFIOptions) (err error) {
	opt.DoubleEncoding = true
	return nil
}

// function designed to set the SSLConnection flag for
// a given LFIOptions object.
func UsingSSL(opt *LFIOptions) (err error) {
	opt.SSLConnection = true
	return nil
}

