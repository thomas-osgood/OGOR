package lfichecker

import "errors"

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

// function designed to add a parameter to the LFI
// testing options. this will take in a param and
// goodval (value that does not fail).
func WithParameter(param string, goodval string) LFIOptsFunc {
	return func(o *LFIOptions) error {
		o.Parameters[param] = goodval
		return nil
	}
}

func WithTarget(targetfile string) LFIOptsFunc {
	return func(o *LFIOptions) error {
		if len(targetfile) < 1 {
			return errors.New("targetfile must be non-zero length string")
		}
		o.TargetFile = targetfile
		return nil
	}
}
